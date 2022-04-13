package ctx

import (
	"errors"
	"fmt"
)

// The UnionFind represents a union–find data type (also known as the disjoint-sets data type).
// It supports the classic union and find operations, along with a count operation that returns
// the total number of sets.
//
// The union–find data type models a collection of sets containing n elements, with each element
// in exactly one set. The elements are named 0 through n–1. Initially, there are n sets, with
// each element in its own set. The canonical element of a set (also known as the root, identifier,
// leader, or set representative) is one distinguished element in the set. Here is a summary of
// the operations:
//
// - Find(p) returns the canonical element of the set containing p. The find operation returns
// the same value for two elements if and only if they are in the same set.
// - Union(p, q) merges the set containing element p with the set containing element q. That is,
// if p and q are in different sets, replace these two sets with a new set that is the union of the two.
// - Count() returns the number of sets.
//
// The canonical element of a set can change only when the set itself changes during a call to union—it
// cannot change during a call to either find or count.
type UnionFind struct {
	parent []int
	rank   []byte
	count  int
}

func NewUnionFind(n int) (*UnionFind, error) {
	if n < 0 {
		return nil, errors.New("argument is illegal")
	}
	uf := UnionFind{
		parent: make([]int, n),
		rank:   make([]byte, n),
		count:  n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
	}
	return &uf, nil
}

func (uf UnionFind) Find(p int) int {
	err := uf.validate(p)
	if err != nil {
		fmt.Println("Find err:", err)
		return -1
	}
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *UnionFind) Union(p, q int) {
	err := uf.validate(p)
	if err != nil {
		fmt.Println("Union err:", err)
		return
	}
	err = uf.validate(q)
	if err != nil {
		fmt.Println("Union err:", err)
		return
	}
	rootP := uf.Find(p)
	rootQ := uf.Find(q)
	if rootP == rootQ {
		return
	}
	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	uf.count--
}

func (uf UnionFind) Connected(p, q int) bool {
	pf, qf := uf.Find(p), uf.Find(q)
	if pf == -1 || qf == -1 {
		return false
	}
	return pf == qf
}

func (uf UnionFind) Count() int {
	return uf.count
}

func (uf UnionFind) validate(p int) error {
	if p < 0 || p >= len(uf.parent) {
		return errors.New(fmt.Sprintf("index %d is not between 0 and %d", p, len(uf.parent)-1))
	}
	return nil
}
