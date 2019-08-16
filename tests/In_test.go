package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

func in(element interface{}, elements interface{}) bool {
	var typs = reflect.TypeOf(elements)
	var vals = reflect.ValueOf(elements)

	// elements not slice
	switch typs.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		return false
	}

	var len = vals.Len()
	for i := 0; i < len; i++ {
		var v = vals.Index(i).Interface()
		if reflect.DeepEqual(element, v) {
			return true
		}
	}

	return false
}

func TestIn(t *testing.T) {
	// string in []string
	var str = "Hello,World!"
	var s_sli = []string{"Hello,World!", "x"}
	fmt.Printf("%q in %s: %t\n", str, s_sli, in(str, s_sli))
	s_sli = []string{"Hello,World", "x"}
	fmt.Printf("%q in %s: %t\n", str, s_sli, in(str, s_sli))

	var i int = 33
	var i_sli = []int{33, 44, 55}
	fmt.Printf("%d in %d: %t\n", i, i_sli, in(i, i_sli))
	i_sli = []int{44, 55}
	fmt.Printf("%d in %d: %t\n", i, i_sli, in(i, i_sli))

	var x = struct {
		int
		string
	}{1, "1"}
	var y = []struct {
		int
		string
	}{struct {
		int
		string
	}{}}
	fmt.Printf("%v in %v: %t\n", x, y, in(x, y))
	y = []struct {
		int
		string
	}{struct {
		int
		string
	}{}, struct {
		int
		string
	}{1, "1"}}
	fmt.Printf("%v in %v: %t\n", x, y, in(x, y))
}
