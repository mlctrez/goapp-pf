package banner

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Variant string

const (
	Default Variant = "default"
	Info    Variant = "info"
	Success Variant = "success"
	Warning Variant = "warning"
	Danger  Variant = "danger"
)

func Variants() []Variant {
	return []Variant{Default, Info, Success, Warning, Danger}
}

func (v Variant) checkVariant() Variant {
	switch v {
	case Info, Success, Warning, Danger, Default:
		return v
	default:
		return Default
	}
}

func (v Variant) class(div app.HTMLDiv) {
	div.Class("pf-c-banner")
	switch v {
	case Info:
		div.Class("pf-m-info")
	case Success:
		div.Class("pf-m-success")
	case Warning:
		div.Class("pf-m-warning")
	case Danger:
		div.Class("pf-m-danger")
	}
}

func (v Variant) toTitle() string {
	vs := string(v.checkVariant())
	return strings.ToUpper(vs[0:1]) + vs[1:]
}
