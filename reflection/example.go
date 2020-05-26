package reflection

import (
	"fmt"
	"reflect"
)

// FindMethod find method of an input reflect.Value
func FindMethod(rv reflect.Value, name string) (reflect.Value, error) {
	var vptr = rv
	if rv.Kind() != reflect.Ptr {
		vptr = reflect.New(rv.Type())
		vptr.Elem().Set(rv)
	}

	method := vptr.MethodByName(name)
	if method.IsValid() {
		return method, nil
	} else {
		return reflect.Value{}, fmt.Errorf("method '%s' not found in '%s'", name, rv.Type().String())
	}
}
