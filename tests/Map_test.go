package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

func UpdateMapValue(m, k, v interface{}) {
	var val = reflect.ValueOf(m)
	var typ = reflect.TypeOf(m)
	if typ.Kind() != reflect.Map {
		return
	}
	var key = reflect.ValueOf(k)
	var value = reflect.ValueOf(v)
	val.SetMapIndex(key, value)
}

func UpdateMapValueMapValue(m, k0, k1, v interface{}) {
	var val = reflect.ValueOf(m)
	var typ = reflect.TypeOf(m)
	if typ.Kind() != reflect.Map {
		return
	}

	var key0 = reflect.ValueOf(k0)
	var va = val.MapIndex(key0)
	if !va.IsValid() {
		va = reflect.MakeMap(typ.Elem())
		val.SetMapIndex(key0, va)
	}

	typ = va.Type()
	if typ.Kind() != reflect.Map {
		return
	}

	var key1 = reflect.ValueOf(k1)
	var value = reflect.ValueOf(v)
	va.SetMapIndex(key1, value)
}

func TestUpdateMapValue(t *testing.T) {
	var map0 = make(map[int]int)
	fmt.Printf("map: %+#v\n", map0)
	UpdateMapValue(map0, 1, 1)
	fmt.Printf("map: %+#v\n", map0)
	var map1 = make(map[int]map[int]int)
	fmt.Printf("map: %+#v\n", map1)
	UpdateMapValueMapValue(map1, 1, 1, 1)
	fmt.Printf("map: %+#v\n", map1)
	var map2 = make(map[int]map[int]int)
	fmt.Printf("map: %+#v\n", map2)
	var inner_map2 = map2[1]
	if inner_map2 == nil {
		inner_map2 = make(map[int]int)
		map2[1] = inner_map2
	}
	fmt.Printf("map: %+#v\n", map2)
	inner_map2[1] = 1
	fmt.Printf("map: %+#v\n", map2)
}
