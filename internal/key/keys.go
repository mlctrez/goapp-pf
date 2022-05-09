package key

import (
	"fmt"
	"reflect"
)

func PackageAndName(i interface{}) string {
	elem := reflect.TypeOf(i).Elem()
	return fmt.Sprintf("%s/%s", elem.PkgPath(), elem.Name())
}
