package avatar

import (
	"encoding/base64"
	"sync"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Demo() app.UI {
	div := app.Div().ID("goapp-pf-demo-avatar")

	var body []app.UI

	section := func(name string, elem ...app.UI) {
		body = append(body, app.H2().Class("pf-c-title pf-m-2xl").Text(name))
		body = append(body, elem...)
	}
	section("Basic", (&Avatar{Src: ImageSrc()}).UI())
	var sv []app.UI
	for _, sizeType := range Sizes() {
		sv = append(sv, app.H3().Class("pf-c-title pf-m-lg").Text(sizeType.name()))
		sv = append(sv, (&Avatar{Src: ImageSrc(), Size: sizeType}).UI())
	}
	section("Size variations", sv...)

	section("Bordered - light", (&Avatar{Src: ImageSrc(), Border: BorderLight}).UI())
	section("Bordered - dark", (&Avatar{Src: ImageSrc(), Border: BorderDark}).UI())

	// /assets/images/pf-c-brand__logo-on-sm.svg

	return div.Body(body...)
}

var once = &sync.Once{}
var imageSrc string

func ImageSrc() string {
	once.Do(func() {
		imageSrc = "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString([]byte(avatarSvg))
	})
	return imageSrc
}

var avatarSvg = `<svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" 
	 xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
	 viewBox="0 0 36 36" style="enable-background:new 0 0 36 36;" xml:space="preserve">
<style type="text/css">
	.st0{fill-rule:evenodd;clip-rule:evenodd;fill:#F0F0F0;}
	.st1{fill-rule:evenodd;clip-rule:evenodd;fill:#D2D2D2;}
	.st2{fill:#B8BBBE;}
	.st3{fill:#D2D2D2;}
</style>
<rect class="st0" width="36" height="36"/>
<path class="st1" d="M17.7,20.1c-3.5,0-6.4-2.9-6.4-6.4s2.9-6.4,6.4-6.4s6.4,2.9,6.4,6.4S21.3,20.1,17.7,20.1z"/>
<path class="st2" d="M13.3,36l0-6.7c-2,0.4-2.9,1.4-3.1,3.5L10.1,36H13.3z"/>
<path class="st3" d="M10.1,36l0.1-3.2c0.2-2.1,1.1-3.1,3.1-3.5l0,6.7h9.4l0-6.7c2,0.4,2.9,1.4,3.1,3.5l0.1,3.2h4.7
	c-0.4-3.9-1.3-9-2.9-11c-1.1-1.4-2.3-2.2-3.5-2.6s-1.8-0.6-6.3-0.6s-6.1,0.7-6.1,0.7c-1.2,0.4-2.4,1.2-3.4,2.6
	C6.7,27,5.8,32.2,5.4,36H10.1z"/>
<path class="st2" d="M25.9,36l-0.1-3.2c-0.2-2.1-1.1-3.1-3.1-3.5l0,6.7H25.9z"/>
</svg>`
