package sort

import (
	"fmt"
	"time"
)

// Heap as a maximum priority queue.
type Heap struct {
	pq []Comparable // Treat index i as index i-1.
}

func (h *Heap) swap(i, j int) {
	h.pq[i-1], h.pq[j-1] = h.pq[j-1], h.pq[i-1]
}

func (h Heap) less(i, j int) bool {
	return h.pq[i-1].CompareTo(h.pq[j-1]) < 0
}

func (h *Heap) sink(k, n int) {
	for k <= n/2 {
		j := 2 * k
		if j < n && h.less(j, j+1) {
			j++
		}
		if !h.less(k, j) {
			break
		}
		h.swap(k, j)
		k = j
	}
}

func (h *Heap) Sort() {
	ts := time.Now()
	// Heap construction
	n := len(h.pq)
	for k := n / 2; k >= 1; k-- {
		h.sink(k, n)
	}
	// Sort down
	for k := n; k > 1; k-- {
		h.swap(1, k)
		h.sink(1, k-1)
	}
	fmt.Println("heap sort time:", time.Since(ts))
}

func (h Heap) Show() {
	n := len(h.pq)
	fmt.Println("is sorted?", h.IsSorted())
	if n > 20 {
		fmt.Println("heap's top 10:", h.pq[:10])
		fmt.Println("heap's final 10:", h.pq[n-10:])
	} else {
		fmt.Println("after heap sort:", h.pq)
	}
}

func (h Heap) IsSorted() bool {
	n := len(h.pq)
	for i := 1; i < n; i++ {
		if h.less(i+1, i) {
			return false
		}
	}
	return true
}
