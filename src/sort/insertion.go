package sort

import (
	"fmt"
	"time"
)

type Insertion struct {
	slice Comparable
}

func (s Insertion) Sort() {
	var n = s.slice.Length()
	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if s.slice.Compare(i, j) {
				s.slice.Swap(i, j)
			}
		}
	}
}

func (s Insertion) ConnSort() {
	var n = s.slice.Length()
	for i := 1; i < n; i++ {
		/*// Note: this section is buggy!
		for j := 0; j < i && s.slice.Compare(j+1, j); j++ {
			s.slice.Swap(j+1, j)
		}*/

		for j := i; j > 0 && s.slice.Compare(j, j-1); j-- {
			s.slice.Swap(j, j-1)
		}
	}
}

func InsertionSort(slice Comparable) {
	ts := time.Now()
	insertion := Insertion{slice: slice}
	insertion.ConnSort()
	fmt.Println("insertion sort time:", time.Since(ts))
}
