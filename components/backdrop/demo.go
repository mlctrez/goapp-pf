package backdrop

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/mlctrez/goapp-pf/internal/ui"
	"github.com/mlctrez/goapp-pf/layouts/bullseye"
)

func DemoBasic() app.UI {
	return (&BackDrop{}).UI()
}

func DemoWithSpinner() app.UI {
	bull := &bullseye.BullsEye{Children: ui.S(
		app.Raw("<svg class=\"pf-c-spinner pf-m-xl\" role=\"progressbar\" viewBox=\"0 0 100 100\" " +
			"aria-label=\"Loading...\"><circle class=\"pf-c-spinner__path\" cx=\"50\" cy=\"50\" " +
			"r=\"45\" fill=\"none\"></circle></svg>"),
	)}

	return (&BackDrop{Children: ui.S(bull.UI())}).UI()
}
