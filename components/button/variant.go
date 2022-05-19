package button

import (
	"fmt"
)

type Variant string

const (
	Primary   Variant = "primary"
	Secondary Variant = "secondary"
	Tertiary  Variant = "tertiary"
	Danger    Variant = "danger"
	Warning   Variant = "warning"
	Link      Variant = "link"
	PLain     Variant = "plain"
	Control   Variant = "control"
)

func Variants() []Variant {
	return []Variant{Primary, Secondary, Tertiary, Danger, Warning, Link, PLain, Control}
}

func (v Variant) checkVariant() Variant {
	switch v {
	case Primary, Secondary, Tertiary, Danger, Warning, Link, PLain, Control:
		return v
	default:
		return Primary
	}
}

func (v Variant) class(btn abs) {
	btn.Class(fmt.Sprintf("pf-m-%s", v))
}
