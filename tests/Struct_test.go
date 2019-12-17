package golangtest

import (
	"fmt"
	"reflect"
	"testing"
)

type Person struct {
	Name   string
	Age    int
	IsMale bool
}

var validColumn = map[reflect.Kind]interface{}{
	reflect.Bool:    nil,
	reflect.Int:     nil,
	reflect.Int8:    nil,
	reflect.Int16:   nil,
	reflect.Int32:   nil,
	reflect.Int64:   nil,
	reflect.Uint:    nil,
	reflect.Uint8:   nil,
	reflect.Uint16:  nil,
	reflect.Uint32:  nil,
	reflect.Uint64:  nil,
	reflect.Float32: nil,
	reflect.Float64: nil,
	reflect.String:  nil,
	reflect.Struct:  nil,
}

var invalidColumn = map[reflect.Kind]interface{}{
	reflect.Invalid:       nil,
	reflect.Uintptr:       nil,
	reflect.Complex64:     nil,
	reflect.Complex128:    nil,
	reflect.Array:         nil,
	reflect.Chan:          nil,
	reflect.Func:          nil,
	reflect.Interface:     nil,
	reflect.Map:           nil,
	reflect.Ptr:           nil,
	reflect.Slice:         nil,
	reflect.UnsafePointer: nil,
}

func getColumnName(valueField reflect.Value, typeField reflect.StructField) (name string, isColumn bool, isStruct bool) {
	if !valueField.CanInterface() {
		return
	}

	var typ = valueField.Type()
	k := typ.Kind()
	_, ok := validColumn[k]
	if !ok {
		return
	}

	if k == reflect.Struct && typ.Name() != "Time" {
		isStruct = true
		return
	}

	isColumn = true
	name = typeField.Name

	return
}

func tableReceivers(val reflect.Value) map[string]interface{} {
	var tr = make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		var columnName, isColumn, isStruct = getColumnName(valueField, typeField)

		if isStruct {
			rs := tableReceivers(valueField)
			for k, v := range rs {
				tr[k] = v
			}
			continue
		}

		if !isColumn {
			continue
		}

		tr[columnName] = valueField.Addr()
	}
	return tr
}

func TestStruct(t *testing.T) {
	var p Person
	var val = reflect.Indirect(reflect.ValueOf(&p))
	var rs = tableReceivers(val)
	fmt.Printf("%+#v\n", rs)
}
