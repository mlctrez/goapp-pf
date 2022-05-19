package flex

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type CssMap map[string]string

type Flex struct {
	Align          CssMap
	AlignContent   CssMap
	AlignItems     CssMap
	AlignSelf      CssMap
	Children       []app.UI
	Component      string
	Direction      CssMap
	Display        CssMap
	Flex           CssMap
	FlexWrap       CssMap
	FullWidth      CssMap
	Grow           CssMap
	JustifyContent CssMap
	Order          CssMap
	Shrink         CssMap
	SpaceItems     CssMap
	Spacer         CssMap
}

func (f *Flex) Render() app.UI {
	return app.Div().Class("pf-l-flex").Body(f.Children...)
}

type Item struct {
	Align     CssMap
	AlignSelf CssMap
	Children  []app.UI
	Component string
	Flex      CssMap
	FullWidth CssMap
	Grow      CssMap
	Order     CssMap
	Shrink    CssMap
	Spacer    CssMap
}

func (i *Item) Render() app.UI {
	return app.Div().Body(i.Children...)
}
