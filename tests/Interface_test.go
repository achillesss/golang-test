package golangtest

import (
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	var i interface{}
	var val = reflect.ValueOf(i)
	println(val.IsValid())
	// true
	println(i == nil)
	var str *string
	// true
	println(str == nil)
	i = str
	// true
	println(i == (*string)(nil))
	// true
	println(i.(*string) == nil)
	// false
	println(i == nil)

	val = reflect.ValueOf(i)
	// true
	println(val.IsNil())
}
