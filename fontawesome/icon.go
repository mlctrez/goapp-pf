package fontawesome

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func Icon(name string) app.HTMLI {
	class := fmt.Sprintf("fas fa-%s", name)
	return app.I().Class(class).Aria("hidden", true)
}
