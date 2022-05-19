package button

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/key"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

type Button struct {
	ID                   string
	AriaLabel            string
	Children             []app.UI
	ClassName            []string
	Component            string
	Icon                 app.HTMLI
	IconPosition         string
	InoperableEvents     []string
	Active               bool
	AriaDisabled         bool
	Block                bool
	Danger               bool
	Disabled             bool
	Inline               bool
	Large                bool
	Loading              bool
	Small                bool
	SpinnerAriaLabel     string
	SpinnerAriaLabeledBy string
	SpinnerAriaValueText string
	TabIndex             *int
	Type                 Type
	Variant              Variant
	OnClick              app.EventHandler
	Href                 string
	Target               string
}

type Type string

const (
	TypeButton Type = "button"
	TypeSubmit Type = "submit"
	TypeReset  Type = "reset"
)

func (t Type) String() string {
	// wasm fails to compile just returning t
	return string(t)
}

func (t Type) checkType() Type {
	switch t {
	case TypeButton, TypeSubmit, TypeReset:
		return t
	default:
		return TypeButton
	}
}

func (b *Button) UI() app.UI {
	return &button{state: *b}
}

func (b *Button) UpdateState(ctx app.Context) {
	ctx.SetState(stateKey(b.ID, "state"), b)
}

type button struct {
	app.Compo
	state Button
}

func (b *Button) init() {
	b.Variant = b.Variant.checkVariant()
	b.Type = b.Type.checkType()
}

func (b *button) OnInit() {
	b.state.init()
}

func (b *button) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(b.state.ID, "state")).Value(&b.state)
}

func (b *button) Render() app.UI {
	return b.state.render()
}

func (b *Button) render() app.UI {

	btn := b.component()
	btn.Class("pf-c-button")
	btn.Type(b.Type.String())
	btn.Role("button")

	b.Variant.class(btn)

	if b.ID != "" {
		btn.ID(b.ID)
	}
	if b.TabIndex != nil {
		btn.TabIndex(*b.TabIndex)
	}

	// optional classes
	btn.ClassIf(b.Small, "pf-m-small")
	btn.ClassIf(b.Large, "pf-m-display-lg")
	btn.ClassIf(b.Inline, "pf-m-inline")
	btn.ClassIf(b.Block, "pf-m-block")
	btn.ClassIf(b.Loading, "pf-m-in-progress")

	if b.Danger {
		Danger.class(btn)
	}

	btn.ClassIf(b.ClassName != nil, b.ClassName...)

	btn.Aria("label", b.AriaLabel)

	if b.AriaDisabled {
		btn.Class("pf-m-aria-disabled")
		btn.Aria("disabled", b.AriaDisabled)
	}
	if b.Disabled {
		btn.Disabled(true)
	}

	var body []app.UI
	body = append(body, b.Children...)

	if b.Icon != nil {
		icnSpan := app.Span().Class("pf-c-button__icon").Body(b.Icon)
		if b.IconPosition == "" || b.IconPosition == "left" {
			body = append([]app.UI{icnSpan.Class("pf-m-start")}, body...)
		} else {
			body = append(body, icnSpan.Class("pf-m-end"))
		}
	}
	if b.Loading {
		progress := app.Span().Class("pf-c-button__progress").Body(b.spinner())
		body = append(ui.S(progress), body...)
	}

	if b.Href != "" {
		btn.Href(b.Href)
	}
	if b.Target != "" {
		btn.Target(b.Target)
	}

	btn.Body(body...)

	if b.OnClick != nil {
		btn.OnClick(b.OnClick)
	}

	return btn.UI()
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&Button{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}

func (b *Button) spinner() app.HTMLSpan {
	span := app.Span().Class("pf-c-spinner pf-m-md")
	span.Aria("role", "progressbar")
	if b.SpinnerAriaLabel != "" {
		span.Aria("label", "Loading...")
	}
	if b.SpinnerAriaLabeledBy != "" {
		span.Aria("labelledby", b.SpinnerAriaLabeledBy)
	}
	if b.SpinnerAriaValueText != "" {
		span.Aria("valuetext", b.SpinnerAriaValueText)
	}
	return span.Body(
		app.Span().Class("pf-c-spinner__clipper"),
		app.Span().Class("pf-c-spinner__lead-ball"),
		app.Span().Class("pf-c-spinner__tail-ball"),
	)

}
