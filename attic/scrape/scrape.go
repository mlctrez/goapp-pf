package scrape

import (
	"bytes"
	"context"
	"fmt"
	"go/format"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/mlctrez/bolt"
	"github.com/pojntfx/html2goapp/pkg/converter"
	"golang.org/x/net/html"
)

type Scrape struct {
	db *bolt.Bolt
}

type ByteString string

func (b ByteString) B() []byte { return []byte(b) }

const (
	DbPath          = "temp/scrape.db"
	Nodes  bolt.Key = "nodes"
)

func New() (*Scrape, error) {
	s := &Scrape{}
	if err := s.initBolt(DbPath); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Scrape) initBolt(path string) (err error) {
	if s.db, err = bolt.New(path); err != nil {
		return err
	}
	return s.db.CreateBuckets(bolt.Keys{Nodes})
}

func (s *Scrape) cleanUp() {
	if s.db != nil {
		if err := s.db.Close(); err != nil {
			log.Println("error closing bolt db", err)
		}
	}
}

func (s *Scrape) Run() (err error) {
	defer s.cleanUp()

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	baseUrl := "https://www.patternfly.org"

	homeRequest := NewRequest(fmt.Sprintf("%s/v4/", baseUrl), "a.pf-c-nav__link")

	if err = s.db.Get(Nodes, homeRequest.storage); err == bolt.ErrValueNotFound {
		if err = chromedp.Run(ctx, selectTask(homeRequest, chromedp.ByQueryAll)); err != nil {
			return err
		}
		err = s.db.Put(Nodes, homeRequest.storage)
	}
	if err != nil {
		return err
	}

	componentNameDuplicationCheck := make(map[string]bool)

	componentNameMap := make(map[string]string)
	var componentUrls []string

	homeRequest.storage.EachNode(func(idx int, node *html.Node) {
		href := nodeAttribute(node, "href")
		if strings.HasPrefix(href, "/v4/components") {
			child := node.FirstChild
			if child.Type == html.TextNode {
				fixed := camelCaseComponentName(child.Data)
				if componentNameDuplicationCheck[fixed] {
					log.Println("duplicate name", fixed)
				}
				componentNameDuplicationCheck[fixed] = true
				componentNameMap[href] = fixed
			}
			componentUrls = append(componentUrls, href)
		}
	})

	var componentRequests []*Request

	for _, componentUrl := range componentUrls {
		compoRequest := NewRequest(baseUrl+componentUrl, "ul.pf-c-tabs__list")
		componentRequests = append(componentRequests, compoRequest)
		if err = s.db.Get(Nodes, compoRequest.storage); err == bolt.ErrValueNotFound {
			if err = chromedp.Run(ctx, selectTask(compoRequest, chromedp.ByQueryAll)); err != nil {
				return err
			}
			err = s.db.Put(Nodes, compoRequest.storage)
		}
	}
	if err != nil {
		return err
	}

	var componentHtmlPages []string
	for _, compoRequest := range componentRequests {
		htmlDemoPage := ""
		reactDemoPage := ""
		compoRequest.storage.EachChildNode(func(idx int, node *html.Node) {
			a := node.FirstChild
			if a != nil && a.FirstChild.Type == html.TextNode {
				href := nodeAttribute(a, "href")
				switch a.FirstChild.Data {
				case "HTML":
					htmlDemoPage = href
				case "React":
					reactDemoPage = href
				}

			}
		})
		if htmlDemoPage != "" {
			componentHtmlPages = append(componentHtmlPages, htmlDemoPage)
		} else if reactDemoPage != "" {
			componentHtmlPages = append(componentHtmlPages, reactDemoPage)
		}
	}

	var compoHtmlPageRequests []*Request
	for _, compoHtmlPage := range componentHtmlPages {
		compoHtmlPageRequest := NewRequest(baseUrl+compoHtmlPage, "div.pf-c-code-editor__controls a")
		compoHtmlPageRequests = append(compoHtmlPageRequests, compoHtmlPageRequest)
		if err = s.db.Get(Nodes, compoHtmlPageRequest.storage); err == bolt.ErrValueNotFound {
			if err = chromedp.Run(ctx, selectTask(compoHtmlPageRequest, chromedp.ByQueryAll)); err != nil {
				return err
			}
			err = s.db.Put(Nodes, compoHtmlPageRequest.storage)
		}
	}
	if err != nil {
		return err
	}

	var compoHtmlDemoPages []string
	for _, request := range compoHtmlPageRequests {
		request.storage.EachNode(func(idx int, node *html.Node) {
			compoHtmlDemoPages = append(compoHtmlDemoPages, nodeAttribute(node, "href"))
		})
	}

	var compoHtmlDemoRequests []*Request

	for _, htmlDemoPageUrl := range compoHtmlDemoPages {
		compoHtmlDemoRequest := NewRequest(baseUrl+htmlDemoPageUrl, "root")
		compoHtmlDemoRequests = append(compoHtmlDemoRequests, compoHtmlDemoRequest)
		if err = s.db.Get(Nodes, compoHtmlDemoRequest.storage); err == bolt.ErrValueNotFound {
			if err = chromedp.Run(ctx, selectTask(compoHtmlDemoRequest, chromedp.ByID)); err != nil {
				return err
			}
			err = s.db.Put(Nodes, compoHtmlDemoRequest.storage)
		}
	}
	if err != nil {
		return err
	}

	// map of go package names to component names
	goAppComponentMap := make(map[string][]string)
	goAppPackageName := "github.com/maxence-charriere/go-app/v9/pkg/app"

	for i, request := range compoHtmlDemoRequests {
		_ = i
		nodes := request.storage.nodes
		if len(nodes) != 1 {
			log.Fatal("length should always be 1 for compoHtmlDemoRequest")
		}
		// drill down to find a class that does not contain ws- prefix
		child := nodes[0].FirstChild
		for child != nil {
			childId := nodeAttribute(child, "id")
			if strings.HasPrefix(childId, "ws-") {
				child = child.FirstChild
				continue
			}
			childClass := nodeAttribute(child, "class")
			if strings.HasPrefix(childClass, "ws-") {
				child = child.FirstChild
				continue
			}
			break
		}
		if child == nil {
			log.Fatal("could not find child")
		}

		if child.NextSibling != nil {
			child = child.Parent
		}

		parse, errParse := url.Parse(request.url)
		if errParse != nil {
			log.Fatal(errParse)
		}

		splitOn := "/html/"
		if strings.Index(parse.Path, splitOn) == -1 {
			splitOn = "/react/"
		}
		split := strings.Split(parse.Path, splitOn)
		if len(split) != 2 {
			log.Fatal("unable to split", parse.Path)
		}

		componentName := componentNameMap[split[0]]
		exampleName := camelCaseComponentName(strings.ReplaceAll(split[1], "-", " "))

		dirName := cleanPackageName(componentName)
		newDir := filepath.Join("components", dirName)
		errMkdir := os.MkdirAll(newDir, 0755)
		if errMkdir != nil {
			log.Fatal(errMkdir)
		}

		out := &bytes.Buffer{}
		errRender := html.Render(out, child)
		if errRender != nil {
			log.Fatal(errRender)
		}

		component, errConvert := converter.ConvertHTMLToComponent(out.String(), goAppPackageName, dirName, exampleName)
		if errConvert != nil {
			log.Fatal(errConvert)
		}

		goAppComponentMap[componentName] = append(goAppComponentMap[componentName], exampleName)

		goFile := filepath.Join(newDir, strings.ToLower(exampleName)+".go")
		errWrite := os.WriteFile(goFile, []byte(component), 0644)
		if errWrite != nil {
			log.Fatal(errWrite)
		}

	}

	buf := &bytes.Buffer{}

	fullPackageName := func(componentName string) string {
		return "github.com/mlctrez/goapp-pf/components/" + cleanPackageName(componentName)
	}

	buf.WriteString("package components\n\n")
	buf.WriteString("import (\n")
	buf.WriteString(fmt.Sprintf("\t%q\n", goAppPackageName))
	for pk := range goAppComponentMap {
		buf.WriteString(fmt.Sprintf("\t%q\n", fullPackageName(pk)))
	}
	buf.WriteString(")\n")

	var componentNames []string
	for componentName := range goAppComponentMap {
		componentNames = append(componentNames, componentName)
	}
	sort.Strings(componentNames)

	for _, componentName := range componentNames {

		exampleNames := goAppComponentMap[componentName]

		packageName := cleanPackageName(componentName)
		buf.WriteString(
			fmt.Sprintf("func %s() map[string]func() app.UI {\n", componentName))
		buf.WriteString("\tresult := make(map[string]func() app.UI)\n")
		for _, name := range exampleNames {
			buf.WriteString(
				fmt.Sprintf("\tresult[%q] = func() app.UI { return &%s.%s{} }\n", name, packageName, name))
		}
		buf.WriteString("\treturn result\n")
		buf.WriteString("}\n\n")

	}

	sort.Strings(componentNames)

	buf.WriteString("func AllComponents() []func() map[string]func() app.UI {\n")
	buf.WriteString("\tvar result []func() map[string]func() app.UI\n")
	for _, name := range componentNames {
		buf.WriteString(fmt.Sprintf("\tresult = append(result, %s)\n", name))
	}
	buf.WriteString("\treturn result\n")
	buf.WriteString("}\n")

	source, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Println(buf.String())
		return err
	}

	err = os.WriteFile("components/factory.go", source, 0644)

	return err
}

func cleanPackageName(componentName string) string {
	packageName := strings.ToLower(componentName)
	if packageName == "select" {
		packageName = "pfselect"
	}
	if packageName == "switch" {
		packageName = "pfswitch"
	}
	return packageName
}

func camelCaseComponentName(name string) string {

	filterDashes := strings.ReplaceAll(name, "-", "")

	filterSpaces := strings.ReplaceAll(filterDashes, "   ", " ")
	filterSpaces = strings.ReplaceAll(filterSpaces, "  ", " ")

	parts := strings.Split(filterSpaces, " ")
	for i, part := range parts {
		part = strings.ToUpper(part[0:1]) + part[1:]
		if i == 0 && part == "Application" {
			part = "App"
		}
		parts[i] = part
	}

	return strings.Join(parts, "")
}

type Request struct {
	url      string
	selector string
	storage  *NodeStorage
}

func NewRequest(url, selector string) *Request {
	key := bolt.Key(fmt.Sprintf("%s %s", url, selector))
	return &Request{
		url:      url,
		selector: selector,
		storage:  &NodeStorage{key: key},
	}
}

func selectTask(request *Request, opts ...chromedp.QueryOption) chromedp.Tasks {
	var nodes []*cdp.Node
	return chromedp.Tasks{
		chromedp.Navigate(request.url),
		chromedp.Nodes(request.selector, &nodes, opts...),
		chromedp.ActionFunc(func(c context.Context) error {
			for _, node := range nodes {
				err := dom.RequestChildNodes(node.NodeID).WithDepth(-1).Do(c)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Sleep(time.Second * 2),
		chromedp.ActionFunc(func(c context.Context) error {
			result := make([]*html.Node, len(nodes))

			for i := 0; i < len(nodes); i++ {
				result[i] = &html.Node{}
				cdpNodeToHtmlNode(nodes[i], result[i])
			}
			request.storage.nodes = result
			fmt.Printf(" %s sel %s length %d\n", request.url, request.selector, len(result))
			return nil
		}),
	}
}
