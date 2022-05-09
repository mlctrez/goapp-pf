package avatar

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Avatar struct {
	Alt       string
	Border    BorderType
	ClassName []string
	Size      SizeType
	Src       string
}

func (a *Avatar) UI() app.UI {
	return &avatar{Avatar: *a}
}

type avatar struct {
	app.Compo
	Avatar
}

func (a *avatar) Render() app.UI {
	img := app.Img().Src(a.Src)

	img.Class("pf-c-avatar")

	if a.Alt != "" {
		img.Alt(a.Alt)
	}
	if a.Border != "" {
		img.Class(a.Border.class())
	}
	if a.Size != "" {
		img.Class(a.Size.class())
	}

	return img
}

func (s SizeType) class() string {
	return fmt.Sprintf("pf-m-%s", s)
}

func (b BorderType) class() string {
	return fmt.Sprintf("pf-m-%s", b)
}
