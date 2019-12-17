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

func initMap(src interface{}) bool {
	var val = reflect.ValueOf(src)
	if val.Kind() != reflect.Ptr {
		if val.Kind() == reflect.Map {
			return true
		}
		return false
	}

	if val.Kind() == reflect.Ptr {
		var elm = val.Elem()
		if elm.Kind() != reflect.Map {
			return false
		}
		elm.Set(reflect.MakeMap(elm.Type()))
	}
	return true
}

func CopyMap(src, dst interface{}) bool {
	if !initMap(dst) {
		return false
	}

	var vals = reflect.ValueOf(src)
	var vald = reflect.Indirect(reflect.ValueOf(dst))
	var typs = vals.Type()
	var typd = vald.Type()

	if typs.Kind() != reflect.Map && typs.Kind() != typd.Kind() {
		return false
	}

	var iter = vals.MapRange()
	for iter.Next() {
		k := iter.Key()
		v := iter.Value()
		t := v.Type()
		if t.Kind() == reflect.Map {
			var km = reflect.MakeMap(t)
			CopyMap(v.Interface(), km.Interface())
			v = km
		}
		vald.SetMapIndex(k, v)
	}
	return true
}

func TestUpdateMapValue(t *testing.T) {
	var nilMap map[int]int
	for k, v := range nilMap {
		println("nil map", k, v)
	}
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

	var src = make(map[int]map[int]int)
	src[1] = map[int]int{1: 1}
	// var dst = make(map[string]string)
	var dst map[int]map[int]int
	fmt.Printf("src: %+v, dst: %+v\n", src, dst)
	// initMap(dst)
	CopyMap(src, &dst)
	fmt.Printf("src: %+v, dst: %+v\n", src, dst)
	// 	dst["y"] = "y"
	// 	fmt.Printf("src: %+v, dst: %+v\n", src, dst)
}
