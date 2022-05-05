package spinner

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type CustomSize struct {
	app.Compo
}

func (c *CustomSize) Render() app.UI {
	return app.Raw("<svg class=\"pf-c-spinner\" role=\"progressbar\" viewBox=\"0 0 100 100\" aria-label=\"Loading...\" style=\"--pf-c-spinner--diameter: 80px;\">\n  <circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" r=\"45\" fill=\"none\"></circle>\n</svg>")
}
