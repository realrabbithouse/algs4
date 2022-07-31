package sort

import (
	"algs4/typ"
)

//
// 思路：逐渐拆分成更小的part，然后再merge，用一个辅组数组将merge的结果写回
//

// IntSliceMergeSort 涉及到赋值，拷贝，无法仅使用三个通用接口实现
type IntSliceMergeSort struct {
	slice typ.IntSlice
}

// merge the first (low to mid) and the second (mid + 1 to high) sorted part into one sorted array.
func (s IntSliceMergeSort) merge(aux typ.IntSlice, low, mid, high int) {
	for i := low; i <= high; i++ {
		aux[i] = s.slice[i]
	}
	i, j := low, mid+1 // 采用双指针进行merge
	for k := low; k <= high; k++ {
		if i > mid {
			s.slice[k] = aux[j]
			j++
		} else if j > high {
			s.slice[k] = aux[i]
			i++
		} else if aux.Compare(i, j) {
			s.slice[k] = aux[i]
			i++
		} else {
			s.slice[k] = aux[j]
			j++
		}
	}
}

func (s IntSliceMergeSort) sortHelper(aux typ.IntSlice, low, high int) {
	// 退出条件
	if high <= low {
		return
	}
	mid := low + (high-low)/2
	s.sortHelper(aux, low, mid)
	s.sortHelper(aux, mid+1, high)
	// sort first, merge second
	s.merge(aux, low, mid, high)
}

func (s IntSliceMergeSort) sort() {
	aux := make(typ.IntSlice, s.slice.Len())
	s.sortHelper(aux, 0, s.slice.Len()-1)
}

func MergeSortInt(slice []int) {
	merge := IntSliceMergeSort{slice}
	merge.sort()
}
