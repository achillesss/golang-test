package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

func useVariables(someVariables interface{}) {
	var val = reflect.ValueOf(someVariables)
	var typ = reflect.TypeOf(someVariables)

	fmt.Printf("use variables: %+#v, %v", someVariables, typ)

	switch typ.Kind() {
	case reflect.Interface, reflect.Ptr:
		fmt.Printf(", element: %v", val.Elem())
		if val.Elem().CanAddr() {
			fmt.Printf(", addr: %x", val.Elem().Addr())
		}
	default:
		var addr = &someVariables
		fmt.Printf(", inner addr: %p", addr)
	}

	println()
}

func TestUseVariables(t *testing.T) {
	var slice = []string{"1", "2"}
	fmt.Printf("variables: %+[1]v, %[1]T, %p\n", slice, &slice)
	useVariables(slice)
	useVariables(&slice)

	var num = 3.1415926
	fmt.Printf("variables: %+[1]v, %[1]T, %p\n", num, &num)
	useVariables(num)
	useVariables(&num)

	var name = "Jason"
	fmt.Printf("variables: %+[1]v, %[1]T, %p\n", name, &name)
	useVariables(name)
	useVariables(&name)

	var alpha = 'Y'
	fmt.Printf("variables: %+[1]v, %[1]T, %p\n", alpha, &alpha)
	useVariables(alpha)
	useVariables(&alpha)
}
