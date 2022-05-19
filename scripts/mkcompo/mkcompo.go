//go:build ignore

package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/mlctrez/goapp-pf/scripts"
)

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr

	flag.Parse()
	if len(flag.Args()) < 2 {
		noErr(fmt.Errorf("Usage: mkcompo <directory> <name>"))
	}

	packageName := flag.Arg(0)
	componentName := flag.Arg(1)
	force := flag.Arg(2) == "force"

	componentDir := filepath.Join("components", packageName)
	fmt.Println(componentDir, componentName)

	componentFile := filepath.Join(componentDir, "component.go")

	if force {
		os.Remove(componentFile)
	}

	_, err := os.Stat(componentFile)
	if err == nil {
		noErr(fmt.Errorf("%s exists, not overwriting it!", componentFile))
	}

	lowerComponent := strings.ToLower(componentName[0:1]) + componentName[1:]

	data := map[string]string{
		"pp": "github.com/mlctrez/goapp-pf/components",
		"pn": packageName,
		"uc": componentName,
		"lc": lowerComponent,
		"rc": lowerComponent[0:1],
	}

	tem := template.Must(template.New("component.go").Parse(tmpl))
	buf := &bytes.Buffer{}
	noErr(tem.Execute(buf, data))

	source, err := format.Source(buf.Bytes())
	noErr(err)

	err = os.WriteFile(componentFile, source, 0644)
	noErr(err)

}

var tmpl = `
package {{.pn}}

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type {{.uc}} struct {
	ID            string
	ClassName     []string
}

func ({{.rc}} *{{.uc}}) UI() app.UI {
	if {{.rc}}.ID == "" {
		{{.rc}}.ID = uuid.NewString()
	}
	return &{{.lc}} { state: *{{.rc}} }
}

func ({{.rc}} *{{.uc}}) UpdateState(ctx app.Context) {
	ctx.SetState(stateKey({{.rc}}.ID, "state"), {{.rc}})
}

func ({{.rc}} *{{.uc}}) onInit() {
	// TODO: any state initialization
}

func ({{.rc}} *{{.uc}}) render() app.UI {
	return app.Div().Text("TODO: ({{.rc}} *{{.uc}}).render()")
}

/*



Everything below here is boilerplate and should not need modification



*/

var _ app.Mounter = (*{{.lc}})(nil)
var _ app.Initializer = (*{{.lc}})(nil)

type {{.lc}} struct {
	app.Compo
	state {{.uc}}
}

func ({{.rc}} *{{.lc}}) OnMount(ctx app.Context) { 
	ctx.ObserveState(stateKey({{.rc}}.state.ID,"state")).Value(&{{.rc}}.state) 
}

func ({{.rc}} *{{.lc}}) OnInit() {
	{{.rc}}.state.onInit()
}

func ({{.rc}} *{{.lc}}) Render() app.UI {
	return {{.rc}}.state.render()
}

func stateKey(id, name string) string {	
	prefix := "{{.pp}}/{{.pn}}/{{.uc}}"
	return fmt.Sprintf("%s/%s/%s", prefix, id, name)
}

`
