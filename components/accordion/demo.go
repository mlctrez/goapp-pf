package accordion

import (
	"math/rand"
	"strings"
	"time"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"gopkg.in/loremipsum.v1"
)

func Demo() *Accordion {
	a := &Accordion{
		ID:               "accordion",
		AriaLabel:        "This is an accordion",
		AsDefinitionList: true,
		DisplaySize:      "default",
		HeadingLevel:     "h3",
		Bordered:         true,
		MultipleExpand:   false,
		Fixed:            false,
	}
	a.Items = buildItems(a.Fixed)
	return a
}

func buildItems(fixed bool) (result []*Item) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	lorem := loremipsum.NewWithSeed(int64(random.Uint32()))
	for i := 0; i < 10; i++ {
		var c []app.UI
		c = append(c, app.Text(lorem.Paragraph()))
		if fixed {
			c = append(c, app.Text(lorem.Paragraph()))
		}
		title := lorem.Word() + " " + lorem.Word()
		title = strings.ToUpper(title[0:1]) + title[1:]
		item := &Item{
			Toggle:  Toggle{Children: []app.UI{app.Text(title)}},
			Content: Content{Children: c},
		}
		result = append(result, item)
	}
	result[5].Toggle.Expanded = true
	return result
}
