//go:build ignore

package main

import (
	"log"

	"github.com/mlctrez/goapp-pf/scrape"
)

func failErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	s, err := scrape.New()
	failErr(err)

	failErr(s.Run())

}
