package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mlctrez/mystace/source"
)

func main() {

	file := "temp/node/node_modules/@patternfly/react-core/src/components/AboutModal/AboutModal.tsx"
	open, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	s, err := source.FromReadCloser(open)
	if err != nil {
		panic(err)
	}

	var tokens []source.Data
	loops := 0
	for loops < 10 {
		loops++
		var peek source.Data
		if peek = s.Peek(1000); peek.Str == "" {
			break
		}

		if start := strings.Index(peek.Str, ";\n"); start > 0 {
			tokens = append(tokens, s.Read(start+2))
			loops = 0
			continue
		}

	}

	for _, token := range tokens {
		fmt.Println(token)
	}

}
