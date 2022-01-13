package sort

import (
	"fmt"
	"time"
)

type Selection struct {
	slice ComparableSlice
}

func (s *Selection) sort() {
	var n = s.slice.Length()
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

func SelectionSort(slice ComparableSlice) {
	ts := time.Now()
	selection := Selection{slice: slice}
	selection.sort()
	fmt.Println("selection sort time:", time.Since(ts))
}
