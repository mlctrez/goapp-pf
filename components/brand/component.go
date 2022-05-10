package brand

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
)

type Brand struct {
	ID        string
	Alt       string
	Children  []app.HTMLSource
	ClassName []string
	heights   map[string]string
	Src       string
	widths    map[string]string
}

func (b *Brand) UI() app.UI {
	return &brand{Brand: *b}
}

type brand struct {
	app.Compo
	Brand
}

func (b *brand) OnMount(ctx app.Context) {
}

func (b *brand) Render() app.UI {
	if b.Children != nil {
		picture := app.Picture().Class("pf-c-brand pf-m-picture")

		// { default?: string; sm?: string; md?: string; lg?: string; xl?: string; '2xl'?: string; }
		// to
		// --pf-c-brand--Width: 40px; --pf-c-brand--Width-on-sm: 60px; --pf-c-brand--Width-on-md: 220px;
		breakPoints := []string{"default", "sm", "md", "lg", "xl", "2xl"}
		for _, point := range breakPoints {
			if b.widths != nil && b.widths[point] != "" {
				k := "--pf-c-brand--Width"
				if point != "default" {
					k = fmt.Sprintf("--pf-c-brand--Width-on-%s", point)
				}
				picture.Style(k, b.widths[point])
			}
			if b.heights != nil && b.heights[point] != "" {
				k := "--pf-c-brand--Height"
				if point != "default" {
					k = fmt.Sprintf("--pf-c-brand--Height-on-%s", point)
				}
				picture.Style(k, b.widths[point])
			}
		}

		var body []app.UI
		for _, child := range b.Children {
			body = append(body, child)
		}
		body = append(body, app.Img().Src(b.Src).Alt(b.Alt))
		return picture.Body(body...)
	} else {
		img := app.Img().Class("pf-c-brand")
		if b.ClassName != nil {
			img.Class(b.ClassName...)
		}
		img.Src(b.Src)
		img.Alt(b.Alt)
		return img
	}
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Brand{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
