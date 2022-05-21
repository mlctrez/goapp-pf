package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/mlctrez/goapp-pf/scripts"
)

func main() {
	defer scripts.Recover()
	noErr := scripts.NoErr

	open, err := os.Open("typescript/kind/kind.go")
	noErr(err)
	defer func() { noErr(open.Close()) }()

	create, err := os.Create("typescript/kind/map.go")
	noErr(err)

	_, err = create.WriteString(`package kind

func (k Kind) StringValue() string {
	switch k {`)
	noErr(err)

	scanner := bufio.NewScanner(open)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, "ExtraKind") {
			break
		}
		if strings.Contains(text, "=") {
			text = strings.Replace(text, "Kind =", "=", 1)
			parts := strings.Split(text, "=")
			name := strings.TrimSpace(parts[0])
			create.WriteString(fmt.Sprintf(`	case %s:
		return %q
`, name, name))
		}
	}
	_, err = create.WriteString(`	default:
		return "<unknown>" 
	}
}`)
	noErr(err)

}
