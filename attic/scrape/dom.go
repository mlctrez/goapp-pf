package scrape

import (
	"bytes"
	"errors"
	"io"
	"strings"

	"github.com/chromedp/cdproto/cdp"
	"github.com/mlctrez/bolt"
	"golang.org/x/net/html"
)

var _ bolt.ValueProvider = (*NodeStorage)(nil)
var _ bolt.ValueReceiver = (*NodeStorage)(nil)

type NodeStorage struct {
	key   bolt.Key
	nodes []*html.Node
}

const NewLine bolt.Key = "\n"

func (n *NodeStorage) Dump(out io.Writer) error {
	for _, node := range n.nodes {
		if err := html.Render(out, node); err != nil {
			return err
		}
		_, err := out.Write(NewLine.B())
		if err != nil {
			return err
		}
	}
	return nil
}

func (n *NodeStorage) Key() bolt.Key {
	return n.key
}

func (n *NodeStorage) Value() ([]byte, error) {
	if n.nodes == nil {
		return nil, errors.New("no value")
	}
	out := &bytes.Buffer{}
	for _, node := range n.nodes {
		if err := html.Render(out, node); err != nil {
			return nil, err
		}
	}
	return out.Bytes(), nil
}

func (n *NodeStorage) SetValue(val []byte) error {
	parse, err := html.Parse(bytes.NewBuffer(val))
	if err != nil {
		return err
	}
	n.nodes = []*html.Node{}
	// strip html/body which is always added by html parse
	node := parse.FirstChild.LastChild.FirstChild
	for node != nil {
		n.nodes = append(n.nodes, node)
		node = node.NextSibling
	}
	return nil
}

func (n *NodeStorage) EachNode(callback func(idx int, node *html.Node)) {
	for i, h := range n.nodes {
		callback(i, h)
	}
}

func (n *NodeStorage) EachChildNode(callback func(idx int, node *html.Node)) {
	for i, h := range n.nodes {
		if i == 0 {
			ci := 0
			child := h.FirstChild
			for child != nil {
				callback(ci, child)
				ci++
				child = child.NextSibling
			}
		}
	}
}

func cdpNodeToHtmlNode(from *cdp.Node, to *html.Node) {

	switch {
	case from.NodeName == "#text":
		to.Type = html.TextNode
		to.Data = from.NodeValue
		return
	default:

		to.Type = html.ElementNode
		to.Data = strings.ToLower(from.NodeName)

		if n := len(from.Attributes); n > 0 {
			for i := 0; i < n; i += 2 {
				to.Attr = append(to.Attr, html.Attribute{
					Key: from.Attributes[i],
					Val: from.Attributes[i+1],
				})
			}
		}
	}

	count := int(from.ChildNodeCount)
	if count == 0 {
		return
	}

	childNodes := make([]*html.Node, count)
	for i := 0; i < count; i++ {
		childNodes[i] = &html.Node{}
	}
	first := 0
	last := count - 1

	to.FirstChild = childNodes[first]
	to.LastChild = childNodes[last]

	for i := 0; i < count; i++ {
		if i > first {
			childNodes[i].PrevSibling = childNodes[i-1]
		}
		if i < last {
			childNodes[i].NextSibling = childNodes[i+1]
		}
	}
	for i, child := range from.Children {
		cdpNodeToHtmlNode(child, childNodes[i])
	}

}

func nodeAttribute(node *html.Node, key string) string {
	for _, attribute := range node.Attr {
		if attribute.Key == key {
			return attribute.Val
		}
	}
	return ""
}
