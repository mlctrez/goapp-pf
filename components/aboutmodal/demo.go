package aboutmodal

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
)

func Demo() *AboutModal {

	brandImage := &BrandImage{Src: "/assets/images/pf_mini_logo_white.svg", Alt: "PatternFly brand logo"}

	trademark := &Trademark{Children: ui.S(app.Text("Trademark and copyright information here"))}

	demoModal := &AboutModal{
		ID:                   "about-modal",
		BrandImage:           brandImage,
		Children:             ui.S(DemoList()),
		CloseButtonAriaLabel: "Close dialog",
		ProductName:          "Product Name",
		Trademark:            trademark,
	}

	//if time.Now().Second()%2 == 0 {
	//	demoModal.Hero = &Hero{BackgroundImageUrl: "/assets/images/pf-logo-small.svg"}
	//}

	return demoModal
}

func DemoList() app.UI {
	return app.Div().Class("pf-c-content").Body(app.Dl().DataSet("pf-content", true).Body(
		app.Dt().DataSet("pf-content", true).Body(app.Text("CFME Version")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("5.5.3.4.20102789036450")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("Cloudforms Version")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("4.1")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("Server Name")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("40DemoMaster")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("User Name")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("Administrator")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("User Role")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("EvmRole-super_administrator")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("Browser Version")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("601.2")),
		app.Dt().DataSet("pf-content", true).Body(app.Text("Browser OS")),
		app.Dd().DataSet("pf-content", true).Body(app.Text("Mac")),
	))
}
