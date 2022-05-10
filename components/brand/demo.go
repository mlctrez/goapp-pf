package brand

import "github.com/maxence-charriere/go-app/v9/pkg/app"

func Demo() app.UI {

	basic := Brand{Src: "/assets/images/pf_logo.svg", Alt: "Patternfly logo"}
	responsive := Brand{Src: "/assets/images/pf-c-brand__logo-base.jpg", Alt: "Patternfly logo"}
	responsive.Children = []app.HTMLSource{
		app.Source().Media("(min-width: 1200px)").SrcSet("/assets/images/pf-c-brand__logo-on-xl.svg"),
		app.Source().Media("(min-width: 992px)").SrcSet("/assets/images/pf-c-brand__logo-on-lg.svg"),
		app.Source().Media("(min-width: 768px)").SrcSet("/assets/images/pf-c-brand__logo-on-md.svg"),
		app.Source().Media("(min-width: 576px)").SrcSet("/assets/images/pf-c-brand__logo-on-sm.svg"),
		app.Source().SrcSet("/assets/images/pf-c-brand__logo.svg"),
	}
	responsive.widths = map[string]string{
		"default": "40px",
		"sm":      "60px",
		"md":      "220px",
	}

	return app.Div().Body(
		app.H4().Text("Basic"),
		basic.UI(),

		app.H4().Text("Responsive"),
		responsive.UI(),
	)
}
