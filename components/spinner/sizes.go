package spinner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Sizes struct {
	app.Compo
}

func (c *Sizes) Render() app.UI {
	return app.Div().
		Class("ws-preview-html pf-u-h-100").
		Body(
			app.Raw("<svg class=\"pf-c-spinner pf-m-sm\" role=\"progressbar\" viewBox=\"0 0 100 100\" aria-label=\"Loading...\">\n    <circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" r=\"45\" fill=\"none\"></circle>\n  </svg>"),
			app.Raw("<svg class=\"pf-c-spinner pf-m-md\" role=\"progressbar\" viewBox=\"0 0 100 100\" aria-label=\"Loading...\">\n    <circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" r=\"45\" fill=\"none\"></circle>\n  </svg>"),
			app.Raw("<svg class=\"pf-c-spinner pf-m-lg\" role=\"progressbar\" viewBox=\"0 0 100 100\" aria-label=\"Loading...\">\n    <circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" r=\"45\" fill=\"none\"></circle>\n  </svg>"),
			app.Raw("<svg class=\"pf-c-spinner pf-m-xl\" role=\"progressbar\" viewBox=\"0 0 100 100\" aria-label=\"Loading...\">\n    <circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" r=\"45\" fill=\"none\"></circle>\n  </svg>"),
		)
}
