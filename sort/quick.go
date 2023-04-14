package sort

import (
	"math/rand"
	"time"

	"algs4/typ"
)

//
// sort(lo, hi)
// mid = partition(lo, hi)
// sort(lo, mid-1)
// sort(mid+1, hi)
//

type IntSliceQuickSort struct {
	slice typ.IntSlice
}

func (q IntSliceQuickSort) partition(lo, hi int) int {
	// 双指针，找到第一个大于v的i，和第一个小于v的j，然后交换它们
	i, j := lo+1, hi
	v := q.slice[lo]
	for {
		for ; q.slice[i] <= v; i++ {
			if i == hi {
				break
			}
		}
		for ; q.slice[j] > v; j-- {
			if j == lo+1 {
				j--
				break
			}
		}
		if i >= j {
			break
		}
		q.slice[i], q.slice[j] = q.slice[j], q.slice[i]
	}
	q.slice[lo], q.slice[j] = q.slice[j], q.slice[lo]
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
	q.sortHelper(0, q.slice.Len()-1)
}

func shuffle(slice []int) {
	rand.Seed(time.Now().UnixNano())
	n := len(slice)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		slice[i], slice[r] = slice[r], slice[i]
	}
}

func QuickSortInt(slice []int) {
	shuffle(slice)
	quick := IntSliceQuickSort{slice}
	quick.sort()
}
