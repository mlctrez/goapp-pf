package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 2 {
		fmt.Println("Usage: makecomponent <directory> <name>")
		return
	}

	packageName := flag.Arg(0)
	componentName := flag.Arg(1)

	componentDir := filepath.Join("components", packageName)
	fmt.Println(componentDir, componentName)

	componentFile := filepath.Join(componentDir, "component.go")
	_, err := os.Stat(componentFile)
	if err == nil {
		fmt.Println(componentFile, "exists, not overwriting it!")
		return
	}

	create, err := os.Create(componentFile)
	if err != nil {
		log.Fatal(err)
	}
	defer create.Close()

	lowerComponent := strings.ToLower(componentName[0:1]) + componentName[1:]

	data := map[string]string{
		"packageName":    packageName,
		"upperComponent": componentName,
		"lowerComponent": lowerComponent,
		"receiverChar":   lowerComponent[0:1],
	}

	tem := template.Must(template.New("component.go").Parse(tmpl))
	err = tem.Execute(create, data)
	if err != nil {
		log.Fatal()
	}

}

var tmpl = `
package {{.packageName}}

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type {{.upperComponent}} struct {
	ID            string
	ClassName     []string
}

func ({{.receiverChar}} *{{.upperComponent}}) UI() app.UI {
	return &{{.lowerComponent}} { {{.upperComponent}}: *{{.receiverChar}} }
}

type {{.lowerComponent}} struct {
	app.Compo
	{{.upperComponent}}
}

func ({{.receiverChar}} *{{.lowerComponent}}) OnMount(ctx app.Context) {
}

func ({{.receiverChar}} *{{.lowerComponent}}) Render() app.UI {
	return app.Div()
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&{{.upperComponent}}{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}

`
