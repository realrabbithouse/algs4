package sort

import (
	"algs4/src/typ"
	"fmt"
	"time"
)

// MinPQ Minimum Priority Queue.
type MinPQ struct {
	n  int
	pq []typ.Comparable
}

func NewMinPQWithCap(initCap int) *MinPQ {
	pq := make([]typ.Comparable, 1, initCap+1)
	pq[0] = nil
	return &MinPQ{
		n:  0,
		pq: pq,
	}
}

func NewMinPQWithKeys(keys []typ.Comparable) *MinPQ {
	ts := time.Now()
	minPQ := NewMinPQWithCap(len(keys))
	for i := range keys {
		minPQ.Insert(keys[i])
	}
	fmt.Println("maximum priority queue construct time:", time.Since(ts))
	return minPQ
}

func (minPQ MinPQ) IsEmpty() bool {
	return minPQ.n == 0
}

func (minPQ MinPQ) Size() int {
	return minPQ.n
}

func (minPQ *MinPQ) Insert(k typ.Comparable) {
	minPQ.pq = append(minPQ.pq, k)
	minPQ.n++
	minPQ.swim(minPQ.n)
}

func (minPQ MinPQ) Min() typ.Comparable {
	if minPQ.IsEmpty() {
		fmt.Println("priority queue underflow")
		return nil
	}
	return minPQ.pq[1]
}

func (minPQ *MinPQ) DelMin() typ.Comparable {
	if minPQ.IsEmpty() {
		fmt.Println("priority queue underflow")
		return nil
	}
	min := minPQ.pq[1]
	minPQ.swap(1, minPQ.n)
	minPQ.pq[minPQ.n] = nil
	minPQ.n--
	minPQ.sink(1)
	if minPQ.n > 0 && minPQ.n == (len(minPQ.pq)-1)/4 {
		fmt.Println("resize from", len(minPQ.pq), "to", len(minPQ.pq)/2, "where n equals", minPQ.n)
		minPQ.resize(len(minPQ.pq) / 2)
	}
	return min
}

func (minPQ *MinPQ) MinK(k int) []typ.Comparable {
	if k > minPQ.n {
		fmt.Println("elements in priority queue not enough")
		return nil
	}
	minK := make([]typ.Comparable, 0, k)
	for i := 0; i < k; i++ {
		minK = append(minK, minPQ.DelMin())
	}
	return minK
}

func (minPQ MinPQ) isMinHeap() bool {
	for i := 1; i < minPQ.n+1; i++ {
		if minPQ.pq[i] == nil {
			fmt.Println("not min heap due to [1~n]", minPQ.n)
			return false
		}
	}
	for i := minPQ.n + 1; i < len(minPQ.pq); i++ {
		if minPQ.pq[i] != nil {
			fmt.Println("not min heap due to [n+1~cap]")
			return false
		}
	}
	if minPQ.pq[0] != nil {
		fmt.Println("not min heap due to [0]")
		return false
	}
	return minPQ.isMinHeapOrdered()
}

func (minPQ MinPQ) isMinHeapOrdered() bool {
	for i := 1; i <= (minPQ.n-1)/2; i++ {
		if minPQ.less(2*i, i) || minPQ.less(2*i+1, i) {
			fmt.Println("not min heap due to order")
			return false
		}
	}
	return true
}

func (minPQ *MinPQ) resize(m int) {
	pq := make([]typ.Comparable, m)
	copy(pq, minPQ.pq)
	minPQ.pq = pq
}

func (minPQ MinPQ) less(i, j int) bool {
	return minPQ.pq[i].CompareTo(minPQ.pq[j]) < 0
}

func (minPQ *MinPQ) swap(i, j int) {
	minPQ.pq[i], minPQ.pq[j] = minPQ.pq[j], minPQ.pq[i]
}

func (minPQ *MinPQ) swim(k int) {
	for k > 1 && minPQ.less(k, k/2) {
		minPQ.swap(k, k/2)
		k /= 2
	}
}

func (minPQ *MinPQ) sink(k int) {
	var j int
	for 2*k <= minPQ.n {
		j = 2 * k
		if j < minPQ.n && minPQ.less(j+1, j) {
			j++
		}
		if !minPQ.less(j, k) {
			break
		}
		minPQ.swap(k, j)
		k = j
	}
}
