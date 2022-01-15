package sort

import (
	"fmt"
	"time"
)

// IntSliceMergeSort 涉及到赋值，拷贝，无法仅使用三个通用接口实现
type IntSliceMergeSort struct {
	IntSlice
}

// merge the first (low to mid) and the second (mid + 1 to high) sorted part into one sorted array.
func (s IntSliceMergeSort) merge(aux IntSlice, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = s.IntSlice[i]
	}
	i, j := low, mid+1
	for k := low; k <= high; k++ {
		if i > mid {
			s.IntSlice[k] = aux[j]
			j++
		} else if j > high {
			s.IntSlice[k] = aux[i]
			i++
		} else if aux.Compare(i, j) {
			s.IntSlice[k] = aux[i]
			i++
		} else {
			s.IntSlice[k] = aux[j]
			j++
		}
	}
}

func (s IntSliceMergeSort) sortHelper(aux IntSlice, low, high int) {
	if high <= low {
		return
	}
	mid := (low + high) / 2
	s.sortHelper(aux, low, mid)
	s.sortHelper(aux, mid+1, high)
	s.merge(aux, low, mid, high)
}

func (s IntSliceMergeSort) sort() {
	aux := s.New(s.Length())
	s.sortHelper(aux.(IntSlice), 0, s.Length()-1)
}

func MergeSort(slice IntSlice) {
	ts := time.Now()
	merge := IntSliceMergeSort{slice}
	merge.sort()
	fmt.Println("merge sort time:", time.Since(ts))
}
