package golangtest

import (
	"fmt"
	"regexp"
	"testing"
)

var reg = regexp.MustCompile(`([^[:upper:]])([[:upper:]])`)

func TestRegexp(t *testing.T) {
	var str = "CurrencyIDs000UserIDs"
	var s = reg.ReplaceAllString(str, "${1}_${2}")
	fmt.Printf("str: %s\n", s)
}
