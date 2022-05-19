package button

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

func (b *Button) component() abs {
	switch b.Component {
	case "a":
		return &anchorAbs{c: app.A()}
	case "span":
		return &spanAbs{c: app.Span()}
	default:
		return &buttonAbs{c: app.Button()}
	}
}

type abs interface {
	ui.UI
	ID(id string)
	TabIndex(t int)
	Class(v ...string)
	Type(t string)
	Role(t string)
	Aria(k string, i any)
	Disabled(d bool)
	Body(elem ...app.UI)
	OnClick(eh app.EventHandler)
	ClassIf(cond bool, v ...string)
	Href(s string)
	Target(s string)
}

type buttonAbs struct {
	c app.HTMLButton
}

func (b *buttonAbs) UI() app.UI                  { return b.c }
func (b *buttonAbs) TabIndex(id int)             { b.c.TabIndex(id) }
func (b *buttonAbs) ID(id string)                { b.c.ID(id) }
func (b *buttonAbs) Class(v ...string)           { b.c.Class(v...) }
func (b *buttonAbs) Type(t string)               { b.c.Type(t) }
func (b *buttonAbs) Role(t string)               { b.c.Role(t) }
func (b *buttonAbs) Aria(k string, i any)        { b.c.Aria(k, i) }
func (b *buttonAbs) Disabled(d bool)             { b.c.Disabled(d) }
func (b *buttonAbs) Body(elem ...app.UI)         { b.c.Body(elem...) }
func (b *buttonAbs) OnClick(eh app.EventHandler) { b.c.OnClick(eh) }
func (b *buttonAbs) Href(t string)               {}
func (b *buttonAbs) Target(t string)             {}
func (b *buttonAbs) ClassIf(cond bool, v ...string) {
	if cond {
		b.c.Class(v...)
	}
}

type anchorAbs struct {
	c app.HTMLA
}

func (a *anchorAbs) UI() app.UI                  { return a.c }
func (a *anchorAbs) TabIndex(id int)             { a.c.TabIndex(id) }
func (a *anchorAbs) ID(id string)                { a.c.ID(id) }
func (a *anchorAbs) Class(v ...string)           { a.c.Class(v...) }
func (a *anchorAbs) Type(t string)               { a.c.Type(t) }
func (a *anchorAbs) Role(t string)               { a.c.Role(t) }
func (a *anchorAbs) Aria(k string, i any)        { a.c.Aria(k, i) }
func (a *anchorAbs) Disabled(d bool)             {}
func (a *anchorAbs) Body(elem ...app.UI)         { a.c.Body(elem...) }
func (a *anchorAbs) OnClick(eh app.EventHandler) { a.c.OnClick(eh) }
func (a *anchorAbs) Href(t string)               { a.c.Href(t) }
func (a *anchorAbs) Target(t string)             { a.c.Target(t) }
func (a *anchorAbs) ClassIf(cond bool, v ...string) {
	if cond {
		a.c.Class(v...)
	}
}

type spanAbs struct {
	c app.HTMLSpan
}

func (s *spanAbs) UI() app.UI                  { return s.c }
func (s *spanAbs) TabIndex(id int)             { s.c.TabIndex(id) }
func (s *spanAbs) ID(id string)                { s.c.ID(id) }
func (s *spanAbs) Class(v ...string)           { s.c.Class(v...) }
func (s *spanAbs) Type(t string)               {}
func (s *spanAbs) Role(t string)               { s.c.Role(t) }
func (s *spanAbs) Aria(k string, i any)        { s.c.Aria(k, i) }
func (s *spanAbs) Disabled(d bool)             {}
func (s *spanAbs) Body(elem ...app.UI)         { s.c.Body(elem...) }
func (s *spanAbs) OnClick(eh app.EventHandler) { s.c.OnClick(eh) }
func (s *spanAbs) Href(t string)               {}
func (s *spanAbs) Target(t string)             {}
func (s *spanAbs) ClassIf(cond bool, v ...string) {
	if cond {
		s.c.Class(v...)
	}
}
