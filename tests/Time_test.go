package golangtest

import (
	"fmt"
	"testing"
	"time"
)

func roundDaily(t time.Time, timezone int) time.Time {
	return t.Add(time.Duration(timezone) * time.Hour).Truncate(time.Hour * 24).Add(-time.Duration(timezone) * time.Hour)
}

func printTime(datetime time.Time) {
	var format = datetime.Format(time.RFC3339)
	fmt.Printf("%s, len: %d\n", format, len(format))
}

func TestTime(t *testing.T) {
	var datetime = time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC)
	printTime(datetime)

	// var zone = time.FixedZone("Beijing", 8*3600)
	var now = time.Now().UTC()
	var n = 24
	for i := 0; i < n; i++ {
		var then = now.Add(time.Hour * time.Duration(i)).UTC()
		printTime(then)
		printTime(roundDaily(then, 8))
		println()
	}
}
