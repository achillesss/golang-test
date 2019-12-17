package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

type XX interface {
	x()
}

type y struct{}

func (y *y) x() {}

func printInterface(src XX) {
	fmt.Printf("src: %T\n", src)
}

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

	var a y
	printInterface(&a)
	var b XX
	printInterface(b)
	b = &a
	printInterface(b)
}
