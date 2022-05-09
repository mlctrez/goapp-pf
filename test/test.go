package main

import (
	"fmt"
	"reflect"

	"github.com/mlctrez/goapp-pf/components/aboutmodal"
)

func main() {
	a := &aboutmodal.AboutModal{}
	fmt.Println(reflect.TypeOf(a).Elem().Name())
}
