package brand

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Responsive struct {
	app.Compo
}

func (c *Responsive) Render() app.UI {
	return app.Picture().
		Class("pf-c-brand pf-m-picture").
		Style("--pf-c-brand--Width", " 40px; --pf-c-brand--Width-on-sm").
		Style(" 60px; --pf-c-brand--Width-on-md", " 220px;").
		Body(
			app.Source().
				Media("(min-width: 1200px)").
				SrcSet("/assets/images/pf-c-brand__logo-on-xl.svg"),
			app.Source().
				Media("(min-width: 992px)").
				SrcSet("/assets/images/pf-c-brand__logo-on-lg.svg"),
			app.Source().
				Media("(min-width: 768px)").
				SrcSet("/assets/images/pf-c-brand__logo-on-md.svg"),
			app.Source().
				Media("(min-width: 576px)").
				SrcSet("/assets/images/pf-c-brand__logo-on-sm.svg"),
			app.Source().
				SrcSet("/assets/images/pf-c-brand__logo.svg"),
			app.Img().
				Src("/assets/images/pf-c-brand__logo-base.jpg").
				Alt("Fallback patternFly default logo"),
		)
}
