package daos

import (
	"math/rand"
	"time"
)

// Reservoir sampling is a family of randomized algorithms for choosing a simple random sample,
// without replacement, of k items from a population of unknown size n in a single pass over the items.

// A Simple (reservoir sampling) and popular but slow algorithm, commonly known as Algorithm R,
// was created by Alan Waterman.
func Simple(stream []int, k int) (r []int) { // k > 0 && k < len(stream)
	n := len(stream)
	if k >= n || k <= 0 {
		return
	}
	r = make([]int, k)
	var i int
	rand.Seed(time.Now().UnixNano())
	for i = 0; i < k; i++ {
		r[i] = stream[i]
	}
	for ; i < n; i++ {
		j := rand.Intn(i + 1)
		if j < k {
			r[j] = stream[i]
		}
	}
	return
}
