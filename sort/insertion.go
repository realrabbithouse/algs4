package sort

import (
	"algs4/typ"
	"fmt"
	"time"
)

// Insertion loop: i = [1 ~ n), j = [i, 1]
type Insertion struct {
	slice typ.ComparableSlice
}

// This sort function is not a real Insertion policy.
func (s Insertion) fakesort() {
	var n = s.slice.Len()
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if s.slice.Compare(i, j) {
				s.slice.Swap(i, j)
			}
		}
	}
}

func (s Insertion) sort() {
	var n = s.slice.Len()
	for i := 1; i < n; i++ {
		/*// Note: this section is buggy!
		for j := 0; j < i && s.slice.Compare(j+1, j); j++ {
			s.slice.Swap(j+1, j)
		}*/

		for j := i; j > 0; j-- {
			if !s.slice.Compare(j-1, j) {
				s.slice.Swap(j, j-1)
			} else {
				break
			}
		}
	}
}

func InsertionSort(slice typ.ComparableSlice) {
	ts := time.Now()
	insertion := Insertion{slice: slice}
	insertion.sort()
	fmt.Println("insertion sort time:", time.Since(ts))
}
