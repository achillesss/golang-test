package golangtest

import (
	"fmt"
	"sort"
	"testing"
)

func UpdateSliceString(sli []string) {
	if len(sli) == 0 {
		return
	}
	sli[0] = "hello"
}

func TestUpdateSliceString(t *testing.T) {
	var slice = []string{"1", "2"}
	fmt.Printf("slice: %+#v\n", slice)
	UpdateSliceString(slice)
	fmt.Printf("slice: %+#v\n", slice)
}

func TestSortStructSlice(t *testing.T) {
	var s = struct {
		sli []int
	}{sli: []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 0}}
	fmt.Printf("struct: %+#v\n", s)
	sort.Slice(s.sli, func(i, j int) bool {
		return s.sli[i] > s.sli[j]
	})
	fmt.Printf("struct: %+#v\n", s)
}
