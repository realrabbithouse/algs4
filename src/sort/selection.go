package sort

import (
	"algs4/src/typ"
	"fmt"
	"time"
)

type Selection struct {
	slice typ.ComparableSlice
}

func (s *Selection) sort() {
	var n = s.slice.Len()
	for i := 0; i < n; i++ {
		var min = i
		for j := i + 1; j < n; j++ {
			if s.slice.Compare(j, min) {
				min = j
			}
		}
		s.slice.Swap(i, min)
	}
}

func SelectionSort(slice typ.ComparableSlice) {
	ts := time.Now()
	selection := Selection{slice: slice}
	selection.sort()
	fmt.Println("selection sort time:", time.Since(ts))
}
