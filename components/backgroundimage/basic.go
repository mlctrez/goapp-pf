package backgroundimage

import app "github.com/maxence-charriere/go-app/v9/pkg/app"

type Basic struct {
	app.Compo
}

func (c *Basic) Render() app.UI {
	return app.Div().
		Class("pf-c-background-image").
		Body(
			app.Raw("<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"pf-c-background-image__filter\" width=\"0\" height=\"0\">\n    <filter id=\"image_overlay\">\n      <feColorMatrix type=\"matrix\" values=\"1 0 0 0 0\n              1 0 0 0 0\n              1 0 0 0 0\n              0 0 0 1 0\"></feColorMatrix>\n      <feComponentTransfer color-interpolation-filters=\"sRGB\" result=\"duotone\">\n        <feFuncR type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncR>\n        <feFuncG type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncG>\n        <feFuncB type=\"table\" tableValues=\"0.086274509803922 0.43921568627451\"></feFuncB>\n        <feFuncA type=\"table\" tableValues=\"0 1\"></feFuncA>\n      </feComponentTransfer>\n    </filter>\n  </svg>"),
		)
}
