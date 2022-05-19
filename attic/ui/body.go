package ui

import (
	"fmt"
	"reflect"
	"runtime"
	"sort"
	"strconv"
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components"
)

var _ app.AppUpdater = (*Body)(nil)
var _ app.Initializer = (*Body)(nil)

type Body struct {
	app.Compo
	namedComponents []NamedComponent
	index           int
}

func (b *Body) OnInit() {
	var namedComponents []NamedComponent
	for _, cf := range components.AllComponents() {
		funcName := runtime.FuncForPC(reflect.ValueOf(cf).Pointer()).Name()
		componentName := funcName[strings.LastIndex(funcName, ".")+1:]
		for name, component := range cf() {
			nc := NamedComponent{Name: componentName, Example: name, Component: component}
			namedComponents = append(namedComponents, nc)
		}
	}
	sort.SliceStable(namedComponents, func(i, j int) bool {
		if namedComponents[i].Name != namedComponents[j].Name {
			return namedComponents[i].Name < namedComponents[j].Name
		}
		return namedComponents[i].Example < namedComponents[j].Example
	})
	b.namedComponents = namedComponents

}

func (b *Body) OnAppUpdate(ctx app.Context) {
	if ctx.AppUpdateAvailable() {
		ctx.Reload()
	}
}

func (b *Body) updateTitle(ctx app.Context) {
	title := b.namedComponents[b.index].Name + " - " + b.namedComponents[b.index].Example
	ctx.Page().SetTitle(title)
}

func (b *Body) saveIndex(ctx app.Context) {
	ctx.SetState("index", b.index, app.Persist)
}

func (b *Body) OnMount(ctx app.Context) {
	b.updateTitle(ctx)
	ctx.GetState("index", &b.index)
	ctx.Handle("navigation", func(context app.Context, action app.Action) {

		firstIndex, previous, next := b.componentIndexMaps()
		current := b.currentComponent()

		switch action.Tags.Get("name") {
		case "NextExample":
			if b.index < (len(b.namedComponents) - 1) {
				b.index++
			}
		case "PreviousExample":
			if b.index > 0 {
				b.index--
			}
		case "PreviousComponent":
			if previous[current.Name] != "" {
				b.index = firstIndex[previous[current.Name]]
			}
		case "NextComponent":
			if next[current.Name] != "" {
				b.index = firstIndex[next[current.Name]]
			}
		case "ChangeIndex":
			app.Log("changeIndex")
			newIndex := action.Tags.Get("index")
			i, err := strconv.Atoi(newIndex)
			if err != nil {
				return
			}
			i--
			if i < 0 {
				i = 0
			}
			if i > (len(b.namedComponents) - 1) {
				i = len(b.namedComponents) - 1
			}
			b.index = i
		}
		b.saveIndex(ctx)
	})
}

type NamedComponent struct {
	Name      string
	Example   string
	Component func() app.UI
}

func (b *Body) Render() app.UI {

	if b.namedComponents == nil {
		return app.Div().Text("loading")
	}

	switch b.currentComponent().Name {
	case "AlertGroup", "Backdrop", "BackgroundImage", "LoginPage", "Masthead", "Page":
		return app.Div().Body(
			b.navigation(),
			b.currentComponent().Component(),
		)
	default:
		return app.Div().Body(
			b.navigation(),
			&Page{Body: []app.UI{b.currentComponent().Component()}},
		)
	}
}

func (b *Body) currentComponent() NamedComponent {
	return b.namedComponents[b.index]
}

func (b Body) componentIndexMaps() (firstIndex map[string]int, previous, next map[string]string) {

	var componentOrder []string
	firstIndex = make(map[string]int)

	for i, component := range b.namedComponents {
		if _, ok := firstIndex[component.Name]; !ok {
			componentOrder = append(componentOrder, component.Name)
			firstIndex[component.Name] = i
		}
	}

	previous = make(map[string]string)
	next = make(map[string]string)
	for i, componentName := range componentOrder {
		if i > 0 {
			previous[componentName] = componentOrder[i-1]
		}
		if i < len(componentOrder)-1 {
			next[componentName] = componentOrder[i+1]
		}
	}

	return firstIndex, previous, next
}

func (b *Body) navigation() app.UI {

	if b.namedComponents == nil {
		return app.Div()
	}

	_, previous, next := b.componentIndexMaps()

	current := b.namedComponents[b.index]
	nav := &Navigation{
		ComponentName:     current.Name,
		ExampleName:       current.Example,
		Current:           fmt.Sprintf("%d", b.index+1),
		Max:               fmt.Sprintf("%d", len(b.namedComponents)),
		PreviousComponent: previous[current.Name] != "",
		PreviousExample:   b.index > 0,
		NextExample:       b.index < (len(b.namedComponents) - 1),
		NextComponent:     next[current.Name] != "",
	}

	styles := map[string]string{
		"position": "fixed",
		"left":     "0",
		"bottom":   "0",
		"z-index":  "500",
	}

	return app.Div().Styles(styles).Body(nav)
}
