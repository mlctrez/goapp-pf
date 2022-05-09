package backgroundimage

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type BackgroundImage struct {
	ImageSrc  string
	ImageMap  ImageSizeMap
	ClassName []string
	Filter    app.UI
}

func (b *BackgroundImage) UI() app.UI {
	return &backgroundImage{BackgroundImage: *b}
}

type ImageSizeMap map[ImageSize]string

func (m ImageSizeMap) ApplyStyles(div app.HTMLDiv) {

	applyStyle := func(d app.HTMLDiv, suffix, url string) {
		d.Style("--pf-c-background-image--BackgroundImage"+suffix, fmt.Sprintf("url(%q)", url))
	}

	for _, size := range ImageSizeOrder() {
		url := m[size]
		switch size {
		case Xs:
			applyStyle(div, "", url)
		case Xs2x:
			applyStyle(div, "-2x", url)
		case Sm:
			applyStyle(div, "-sm", url)
		case Sm2x:
			applyStyle(div, "-sm-2x", url)
		case Lg:
			applyStyle(div, "-lg", url)
		}
	}

}

type ImageSize string

const (
	Xs   ImageSize = "xs"
	Xs2x ImageSize = "xs2x"
	Sm   ImageSize = "sm"
	Sm2x ImageSize = "sm2x"
	Lg   ImageSize = "lg"
)

type backgroundImage struct {
	app.Compo
	BackgroundImage
}

func (b *backgroundImage) Render() app.UI {
	if b.Filter == nil {
		b.Filter = app.Raw(DefaultFilter)
	}
	if b.ImageSrc != "" && b.ImageMap == nil {
		b.ImageMap = make(ImageSizeMap)
		for _, size := range ImageSizeOrder() {
			b.ImageMap[size] = b.ImageSrc
		}
	}
	div := app.Div().Class("pf-c-background-image").Class(b.ClassName...)
	if b.ImageMap != nil {
		b.ImageMap.ApplyStyles(div)
	}

	div.Body(b.Filter)
	return div
}

func ImageSizeOrder() []ImageSize {
	return []ImageSize{Xs, Xs2x, Sm, Sm2x, Lg}
}

var DefaultFilter = `<svg
    xmlns="http://www.w3.org/2000/svg"
    class="pf-c-background-image__filter"
    width="0"
    height="0"
  >
    <filter id="image_overlay">
      <feColorMatrix
        type="matrix"
        values="1 0 0 0 0
              1 0 0 0 0
              1 0 0 0 0
              0 0 0 1 0"
      />
      <feComponentTransfer color-interpolation-filters="sRGB" result="duotone">
        <feFuncR type="table" tableValues="0.086274509803922 0.43921568627451" />
        <feFuncG type="table" tableValues="0.086274509803922 0.43921568627451" />
        <feFuncB type="table" tableValues="0.086274509803922 0.43921568627451" />
        <feFuncA type="table" tableValues="0 1" />
      </feComponentTransfer>
    </filter>
  </svg>`
