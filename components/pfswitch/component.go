package pfswitch

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	fa "github.com/mlctrez/goapp-pf/fontawesome"
)

type PfSwitch struct {
	ID        string
	AriaLabel string
	ClassName []string
	CheckIcon bool
	LabelId   string
	Checked   bool
	Disabled  bool
	Reversed  bool
	Label     app.UI
	LabelOff  app.UI
	OnChange  app.EventHandler
}

func (p *PfSwitch) UI() app.UI {
	p.onInit()
	return &pfSwitch{state: *p}
}

func (p *PfSwitch) UpdateState(ctx app.Context) {
	ctx.SetState(stateKey(p.ID, "state"), p)
}

func (p *PfSwitch) onInit() {
	if p.ID == "" {
		p.ID = uuid.NewString()
	}
	if p.LabelId == "" {
		p.LabelId = p.ID + "-input"
	}
}

func (p *PfSwitch) render() app.UI {

	label := app.Label().Class("pf-c-switch").For(p.LabelId)
	if p.Reversed {
		label.Class("pf-m-reverse")
	}

	input := app.Input().ID(p.LabelId).Class("pf-c-switch__input").Type("checkbox")
	input.Checked(p.Checked).Disabled(p.Disabled)

	if p.OnChange == nil {
		input.OnClick(func(ctx app.Context, e app.Event) { e.PreventDefault() })
	} else {
		input.OnChange(func(ctx app.Context, e app.Event) {
			p.Checked = ctx.JSSrc().Get("checked").Bool()
			p.UpdateState(ctx)
			p.OnChange(ctx, e)
		})
	}
	if p.AriaLabel != "" {
		input.Aria("label", p.AriaLabel)
	}
	if p.Checked && p.Label != nil {
		input.Aria("labelledby", p.LabelId+"-on")
	} else if p.LabelOff != nil {
		input.Aria("labelledby", p.LabelId+"-off")
	}

	toggle := app.Span().Class("pf-c-switch__toggle")
	if p.CheckIcon {
		toggle.Body(app.Span().Class("pf-c-switch__toggle-icon").Body(fa.Icon("check")))
	}

	body := []app.UI{input, toggle}

	if p.Label != nil {
		body = append(body, p.label("-on", p.Label))
	}
	if p.LabelOff != nil {
		body = append(body, p.label("-off", p.LabelOff))
	}

	return label.Body(body...)
}

func (p *PfSwitch) label(suffix string, ui app.UI) app.UI {
	style := app.Style().Class("pf-c-switch__label").Class("pf-m" + suffix)
	return style.ID(p.LabelId + suffix).Body(ui)
}

/*



Everything below here is boilerplate and should not need modification



*/

var _ app.Mounter = (*pfSwitch)(nil)
var _ app.Initializer = (*pfSwitch)(nil)

type pfSwitch struct {
	app.Compo
	state PfSwitch
}

func (p *pfSwitch) OnMount(ctx app.Context) {
	ctx.ObserveState(stateKey(p.state.ID, "state")).Value(&p.state)
}

func (p *pfSwitch) OnInit() {
	p.state.onInit()
}

func (p *pfSwitch) Render() app.UI {
	return p.state.render()
}

func stateKey(id, name string) string {
	prefix := "github.com/mlctrez/goapp-pf/components/pfswitch/PfSwitch"
	return fmt.Sprintf("%s/%s/%s", prefix, id, name)
}
