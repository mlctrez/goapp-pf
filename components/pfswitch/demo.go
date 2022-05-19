package pfswitch

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Demo() app.UI {

	onChangeDoesNothing := func(ctx app.Context, e app.Event) {}

	basic := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Checked: true, OnChange: onChangeDoesNothing}

	reversedLayout := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Reversed: true, Checked: true, OnChange: onChangeDoesNothing}

	withoutLabel := &PfSwitch{OnChange: onChangeDoesNothing}

	checkedWithLabel := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		CheckIcon: true, Checked: true, OnChange: onChangeDoesNothing}

	disabledOne := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Disabled: true, Checked: true}
	disabledTwo := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Disabled: true, Checked: false}
	disabledThree := &PfSwitch{Disabled: true, CheckIcon: true, Checked: true}
	disabledFour := &PfSwitch{Disabled: true, CheckIcon: true}

	uncontrolledOne := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Checked: true}
	uncontrolledTwo := &PfSwitch{Label: app.Text("Message when on"), LabelOff: app.Text("Message when off"),
		Checked: false}
	uncontrolledThree := &PfSwitch{CheckIcon: true, Checked: true}
	uncontrolledFour := &PfSwitch{CheckIcon: true}

	header := func(name string) app.UI {
		return app.H3().Text(name).Style("padding", "20px 0 20px 0")
	}

	return app.Div().Body(
		header("Basic"), basic.UI(),
		header("Reversed Layout"), reversedLayout.UI(),
		header("Without label"), withoutLabel.UI(),
		header("Checked with label"), checkedWithLabel.UI(),
		header("Disabled"), app.Div().Body(
			disabledOne.UI(), app.Br(),
			disabledTwo.UI(), app.Br(),
			disabledThree.UI(), app.Br(),
			disabledFour.UI(), app.Br(),
		),
		header("Uncontrolled"), app.Div().Body(
			uncontrolledOne.UI(), app.Br(),
			uncontrolledTwo.UI(), app.Br(),
			uncontrolledThree.UI(), app.Br(),
			uncontrolledFour.UI(), app.Br(),
		),
	)

}
