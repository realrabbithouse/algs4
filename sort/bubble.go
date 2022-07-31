package sort

import (
	"algs4/typ"
)

type Bubble struct {
	slice typ.ComparableSlice
}

func (s Bubble) sort() {
	var n = s.slice.Len()
	for i := n; i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if !s.slice.Compare(j, j+1) {
				s.slice.Swap(j, j+1)
			}
		}
	}
}

func BubbleSort(slice typ.ComparableSlice) {
	bubble := Bubble{slice: slice}
	bubble.sort()
}
