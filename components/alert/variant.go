package alert

import (
	"strings"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
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

func (v Variant) icon() app.HTMLI {
	switch v {
	case Info:
		return fa.Icon("info-circle")
	case Success:
		return fa.Icon("check-circle")
	case Warning:
		return fa.Icon("exclamation-triangle")
	case Danger:
		return fa.Icon("exclamation-circle")
	case Default:
		return fa.Icon("bell")
	default:
		return fa.Icon("bell")
	}
}

func (v Variant) toTitle() string {
	vs := string(v.checkVariant())
	return strings.ToUpper(vs[0:1]) + vs[1:]
}
