package sort

import (
	"fmt"
	"time"
)

type MaxPQ struct {
	n  int
	pq []Comparable
}

func NewMaxPQ() *MaxPQ {
	return &MaxPQ{
		n:  0,
		pq: []Comparable{nil},
	}
}

func NewMaxPQWithCap(initCap int) *MaxPQ {
	pq := make([]Comparable, 1, initCap+1)
	pq[0] = nil
	return &MaxPQ{
		n:  0,
		pq: pq,
	}
}

func NewMaxPQWithKeys(keys []Comparable) *MaxPQ {
	ts := time.Now()
	maxPQ := NewMaxPQWithCap(len(keys))
	for i := range keys {
		maxPQ.Insert(keys[i])
	}
	fmt.Println("maximum priority queue construct time:", time.Since(ts))
	return maxPQ
}

func (q *MaxPQ) Insert(v Comparable) {
	q.pq = append(q.pq, v)
	q.n++
	q.swim(q.n)
}

func (q *MaxPQ) Max() Comparable {
	if q.n == 0 {
		panic("priority queue underflow")
	}
	return q.pq[1]
}

func (q *MaxPQ) DelMax() Comparable {
	if q.n == 0 {
		panic("priority queue underflow")
	}
	max := q.pq[1]
	q.swap(1, q.n)
	q.pq[q.n] = nil
	q.n--
	q.sink(1)
	if q.n > 0 && q.n == (cap(q.pq)-1)/4 {
		q.pq = q.pq[:(cap(q.pq)-1)/2]
	}
	return max
}

func (q MaxPQ) IsEmpty() bool {
	return q.n == 0
}

func (q MaxPQ) Size() int {
	return q.n
}

func (q *MaxPQ) swap(i, j int) {
	q.pq[i], q.pq[j] = q.pq[j], q.pq[i]
}

func (q *MaxPQ) less(i, j int) bool {
	return q.pq[i].CompareTo(q.pq[j]) < 0
}

func (q *MaxPQ) swim(k int) {
	for k > 1 && q.less(k/2, k) {
		q.swap(k/2, k)
		k = k / 2
	}
}

func (q *MaxPQ) sink(k int) {
	for 2*k <= q.n {
		j := 2 * k
		if j < q.n && q.less(j, j+1) {
			j++
		}
		if !q.less(k, j) {
			break
		}
		q.swap(k, j)
		k = j
	}
}

func (q MaxPQ) isMaxHeap() bool {
	for i := 1; i < q.n+1; i++ {
		if q.pq[i] == nil {
			fmt.Println("not max heap due to [1~n]")
			return false
		}
	}
	for i := q.n + 1; i < len(q.pq); i++ {
		if q.pq[i] != nil {
			fmt.Println("not max heap due to [n+1~cap]")
			return false
		}
	}
	if q.pq[0] != nil {
		fmt.Println("not max heap due to [0]")
		return false
	}
	return q.isMaxHeapOrdered()
}

func (q MaxPQ) isMaxHeapOrdered() bool {
	for i := 1; i <= (q.n-1)/2; i++ {
		if q.less(i, 2*i) || q.less(i, 2*i+1) {
			fmt.Println("not max heap due to order")
			return false
		}
	}
	return true
}

func (q *MaxPQ) TopK(k int) []Comparable {
	if k > q.n {
		fmt.Println("elements in priority queue not enough")
		return nil
	}
	topK := make([]Comparable, 0, k)
	for i := 0; i < k; i++ {
		topK = append(topK, q.DelMax())
	}

	return topK
}
