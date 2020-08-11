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

func removeElementFromArray(src []string, index int) []string {
	var x = make([]string, 0, cap(src))
	x = append(x, src...)

	var head = x[:index]

	if index == len(x)-1 {
		return head
	}

	var tail = x[index+1:]
	return append(head, tail...)
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

	var src = []string{"1", "2", "3", "4"}
	for i := range src {
		var y = removeElementFromArray(src, i)
		fmt.Printf("y: %v\n", y)
	}
}
