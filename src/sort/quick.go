package sort

import (
	"fmt"
	"time"
)

type IntSliceQuickSort struct {
	IntSlice
}

func (q IntSliceQuickSort) partition(lo, hi int) int {
	//fmt.Println("before partition:", q.IntSlice)
	// 双指针，找到第一个大于v的i，和第一个小于v的j，然后交换它们
	i, j := lo+1, hi
	v := q.IntSlice[lo]
	for {
		for ; q.IntSlice[i] <= v; i++ {
			if i == hi {
				break
			}
		}
		for ; q.IntSlice[j] > v; j-- {
			if j == lo+1 {
				j--
				break
			}
		}
		if i >= j {
			break
		}
		q.IntSlice[i], q.IntSlice[j] = q.IntSlice[j], q.IntSlice[i]
	}
	// s[j] < s[lo], s[>j] > s[lo]
	q.IntSlice[lo], q.IntSlice[j] = q.IntSlice[j], q.IntSlice[lo]
	//fmt.Println("after partition:", q.IntSlice, "seed:", j)
	return j
}

func (q IntSliceQuickSort) sortHelper(lo, hi int) {
	if lo >= hi {
		return
	}
	j := q.partition(lo, hi)
	q.sortHelper(lo, j-1)
	q.sortHelper(j+1, hi)
}

func (q IntSliceQuickSort) sort() {
	q.sortHelper(0, q.Len()-1)
}

func QuickSortInt(slice []int) {
	ts := time.Now()
	quick := IntSliceQuickSort{slice}
	quick.sort()
	fmt.Println("quick sort time:", time.Since(ts))
}
