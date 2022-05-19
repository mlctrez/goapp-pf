package aboutmodal

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/components/backdrop"
	"github.com/mlctrez/goapp-pf/components/button"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/key"
	"github.com/mlctrez/goapp-pf/internal/ui"
	"github.com/mlctrez/goapp-pf/layouts/bullseye"
)

// AboutModal renders an about modal dialog wrapped with backdrop and bullseye.
type AboutModal struct {
	ID                   string
	BrandImage           *BrandImage
	Children             []app.UI
	Hero                 *Hero
	CloseButtonAriaLabel string
	ProductName          string
	Trademark            *Trademark
	Open                 bool
}

func (a *AboutModal) UI() app.UI {
	return &aboutModal{state: *a}
}

func (a *AboutModal) UpdateState(ctx app.Context) {
	ctx.SetState(stateKey(a.ID, "state"), a)
}

type BrandImage struct {
	Src string
	Alt string
}

func (b *BrandImage) UI() app.UI {
	div := app.Div()
	if b != nil {
		img := app.Img().Class("pf-c-about-modal-box__brand-image").Src(b.Src).Alt(b.Alt)
		div.Class("pf-c-about-modal-box__brand").Body(img)
	}
	return div
}

type Trademark struct {
	Children []app.UI
}

func (t *Trademark) UI() app.UI {
	if t != nil {
		return app.Div().Class("pf-c-about-modal-box__strapline").Body(t.Children...)
	}
	return app.Div()
}

type Hero struct {
	BackgroundImageUrl string
}

func (h *Hero) UI() app.UI {
	hero := app.Div().Class("pf-c-about-modal-box__hero")
	if h != nil {
		value := fmt.Sprintf("url(%s)", h.BackgroundImageUrl)
		hero.Style("--pf-c-about-modal-box__hero--sm--BackgroundImage", value)
	}
	return hero
}

type aboutModal struct {
	app.Compo
	state AboutModal
}

func (a *aboutModal) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(a.state.ID, "state")).Value(&a.state)
}

func (a *aboutModal) Render() app.UI {
	div := app.Div().ID(a.state.ID)
	if a.state.Open {
		bullsEye := &bullseye.BullsEye{Children: ui.S(a.aboutModalBox())}
		backDrop := &backdrop.BackDrop{Children: ui.ToAppUI(bullsEye)}
		return div.Body(backDrop.UI())
	}
	return div
}

func (a *aboutModal) aboutModalBox() app.UI {

	closeButton := &button.Button{
		Children:  ui.S(fa.Icon("times")),
		AriaLabel: a.state.CloseButtonAriaLabel,
		OnClick:   func(ctx app.Context, e app.Event) { a.state.Open = false },
		Variant:   button.PLain,
	}

	return app.Div().Class("pf-c-about-modal-box").Aria("role", "dialog").
		Aria("modal", true).Aria("labelledby", a.state.ID+"-title").Body(

		a.state.BrandImage.UI(),

		app.Div().Class("pf-c-about-modal-box__close").Body(closeButton.UI()),

		app.Div().Class("pf-c-about-modal-box__header").Body(
			app.H1().Class("pf-c-title pf-m-4xl").
				ID(a.state.ID+"-title").Text(a.state.ProductName)),

		a.state.Hero.UI(),

		app.Div().Class("pf-c-about-modal-box__content").Body(
			app.Div().Class("pf-c-about-modal-box__body").Body(a.state.Children...),
			a.state.Trademark.UI()),
	)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&AboutModal{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
