package sort

import (
	"algs4/typ"
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
	selection := Selection{slice: slice}
	selection.sort()
}
