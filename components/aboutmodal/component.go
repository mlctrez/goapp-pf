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
// https://www.patternfly.org/v4/components/about-modal/#props
type AboutModal struct {
	ID                   string
	BrandImage           *BrandImage
	Children             []app.UI
	Hero                 *Hero
	CloseButtonAriaLabel string
	ProductName          string
	Trademark            *Trademark
}

func (a *AboutModal) UI() app.UI {
	return &aboutModal{AboutModal: *a}
}

func (a *AboutModal) Open(ctx app.Context, _ app.Event) {
	ctx.SetState(stateKey(a.ID, "open"), true)
}

func (a *AboutModal) Close(ctx app.Context, _ app.Event) {
	ctx.SetState(stateKey(a.ID, "open"), false)
}

type BrandImage struct {
	Src string
	Alt string
}

func (b *BrandImage) UI() app.UI {
	div := app.Div()
	if b != nil {
		img := app.Img().Class("pf-c-about-modal-box__brand-image")
		img.Src(b.Src).Alt(b.Alt)

		div.Class("pf-c-about-modal-box__brand").Body(img)
	}
	return div
}

type Trademark struct {
	Children []app.UI
}

func (b *Trademark) UI() app.UI {
	div := app.Div().Class("pf-c-about-modal-box__strapline")
	if b != nil {
		div.Body(b.Children...)
	}
	return div
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
	AboutModal
	open bool
}

func (a *aboutModal) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(a.ID, "open")).Value(&a.open)
}

func (a *aboutModal) Render() app.UI {
	div := app.Div().ID(a.ID)
	if a.open {
		bullsEye := &bullseye.BullsEye{Children: ui.S(a.aboutModalBox())}
		backDrop := &backdrop.BackDrop{Children: ui.ToAppUI(bullsEye)}
		div.Body(backDrop.UI())
	}
	return div
}

func (a *aboutModal) aboutModalBox() app.UI {

	return app.Div().Class("pf-c-about-modal-box").Aria("role", "dialog").
		Aria("modal", true).Aria("labelledby", a.ID+"-title").Body(

		a.BrandImage.UI(),

		app.Div().Class("pf-c-about-modal-box__close").Body(
			button.Plain(fa.Icon("times")).OnClick(a.Close).
				Aria("label", a.CloseButtonAriaLabel)),

		app.Div().Class("pf-c-about-modal-box__header").Body(
			app.H1().Class("pf-c-title pf-m-4xl").
				ID(a.ID+"-title").Text(a.ProductName)),

		a.Hero.UI(),

		app.Div().Class("pf-c-about-modal-box__content").Body(
			app.Div().Class("pf-c-about-modal-box__body").Body(a.Children...),
			a.Trademark.UI()),
	)
}

var packageAndName = ""

func init() {
	packageAndName = key.PackageAndName(&AboutModal{})
}

func stateKey(componentId, name string) string {
	return fmt.Sprintf("%s/%s/%s", packageAndName, componentId, name)
}
