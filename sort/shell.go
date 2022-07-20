package sort

import (
	"algs4/typ"
	"fmt"
	"time"
)

// Shell is an optimized version of insertion sort.
type Shell struct {
	slice typ.ComparableSlice
}

func (s Shell) sort() {
	n := s.slice.Len()
	for i := n / 3; i > 0; i /= 3 {
		s.hGapSort(i)
	}
}

func (s Shell) hGapSort(h int) {
	n := s.slice.Len()
	for i := h; i < n; i++ {
		for j := i; j >= h; j = j - h {
			if !s.slice.Compare(j-h, j) {
				s.slice.Swap(j-h, j)
			} else {
				break
			}
		}
	}
	//fmt.Printf("after gap %d sort: %v\n", h, s.slice)
}

func ShellSort(slice typ.ComparableSlice) {
	ts := time.Now()
	shell := Shell{slice: slice}
	shell.sort()
	fmt.Println("Shell sort time:", time.Since(ts))
}
