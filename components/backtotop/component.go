package backtotop

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

type BackToTop struct {
	ClassName          []string
	AlwaysVisible      bool
	ScrollableSelector string
	Title              string
}

func (b *BackToTop) UI() app.UI {
	return &backToTop{BackToTop: *b}
}

type backToTop struct {
	app.Compo
	BackToTop
	visible      bool
	scrollTarget app.Value
	scrollFunc   app.Func
}

func (b *backToTop) OnDismount() {
	if b.scrollTarget != nil {
		if b.scrollTarget.Truthy() && b.scrollFunc != nil {
			app.Log("removeEventListener")
			b.scrollTarget.Call("removeEventListener", "scroll", b.scrollFunc)
		}
	}
	if b.scrollFunc != nil {
		app.Log("scrollFunc.Release()")
		b.scrollFunc.Release()
	}
}

func (b *backToTop) OnMount(ctx app.Context) {
	if b.ScrollableSelector != "" {
		if elem, err := ui.QuerySelector(b.ScrollableSelector); err != nil {
			app.Log("warning, query did not match element:" + b.ScrollableSelector)
		} else {
			b.scrollTarget = elem
		}
	} else {
		b.scrollTarget = app.Window()
	}
	b.scrollFunc = app.FuncOf(b.onScrollEvent)
	b.scrollTarget.Call("addEventListener", "scroll", b.scrollFunc)

}

func (b *backToTop) onScrollEvent(this app.Value, args []app.Value) interface{} {
	scrolled := this.Get("scrollY")
	if scrolled.IsUndefined() {
		scrolled = this.Get("scrollTop")
	}
	if !b.AlwaysVisible {
		before := b.visible
		b.visible = scrolled.Int() > 400
		if b.visible != before {
			b.Update()
		}
	}
	return nil
}

func (b *backToTop) buttonClicked(ctx app.Context, e app.Event) {
	if b.scrollTarget.Truthy() {
		b.scrollTarget.Call("scrollTo", map[string]interface{}{"top": 0, "behavior": "smooth"})
	}

}

func (b *backToTop) Render() app.UI {
	div := app.Div()
	if b.AlwaysVisible || b.visible {
		div.Class("pf-c-back-to-top").Body(

			app.Button().Aria("disabled", "false").OnClick(b.buttonClicked).
				Class("pf-c-button pf-m-primary").Type("button").Body(

				app.Text(b.Title),
				app.Span().Class("pf-c-button__icon pf-m-end").Body(fa.Icon("angle-up")),
			),
		)
	}

	return div
}
