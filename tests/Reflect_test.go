package golangtest

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

func dateTimeStr(t time.Time) string {
	return fmt.Sprintf("%q", t.Format(timeFormat))
}

func defaultSQLValue(column string) string {
	if column == "" {
		return "DEFAULT"
	}

	return fmt.Sprintf("DEFAULT(%s)", column)
}

func fmtToSQLValue(src interface{}, column string) interface{} {
	srcValue := reflect.ValueOf(src)
	srcType := reflect.TypeOf(src)

	if srcType.Kind() == reflect.Ptr {
		if srcValue.IsNil() {
			return "NULL"
		}
		srcValue = srcValue.Elem()
		srcType = srcType.Elem()
	}

	zeroValue := reflect.Zero(srcType)
	if reflect.DeepEqual(zeroValue.Interface(), srcValue.Interface()) {
		return defaultSQLValue(column)
	}

	switch srcType.Kind() {
	case reflect.String:
		return fmt.Sprintf("%q", srcValue)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128, reflect.Bool:
		return srcValue

	default:
		if srcType.Name() == "Time" {
			return dateTimeStr(srcValue.Interface().(time.Time).UTC())
		}

		fmt.Printf("bad type: %s\n", srcType.Name())
		return src
	}
}

type X struct {
	x *bool
}

type Y struct {
	X
}

func TestReflect(t *testing.T) {
	var a Y
	var tt = true
	a.x = &tt
	fmtToSQLValue(a.x, "is_end")
}
