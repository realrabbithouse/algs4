package graphapp

import (
	"algs4/src/graph"
	"errors"
	"fmt"
)

// Topological 拓扑排序
type Topological struct {
	order []graph.ID
	rank  []int
}

func NewTopological(digraph *graph.Digraph) (*Topological, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	t := Topological{
		rank: make([]int, digraph.V),
	}
	cycle, _ := NewDirectedCycle(digraph)
	if !cycle.HasCycle() {
		dfo, _ := NewDepthFirstOrder(digraph)
		t.order = dfo.ReversePostOrder()
		for i, id := range t.order {
			t.rank[id] = i
		}
	}
	return &t, nil
}

func (t *Topological) HasOrder() bool {
	return t.order != nil
}

func (t *Topological) Order() []graph.ID {
	return t.order
}

func (t *Topological) Rank(v int) int {
	err := t.validateVertex(v)
	if err != nil {
		fmt.Println("Rank err:", err)
		return -1
	}
	return t.rank[v]
}

func (t *Topological) validateVertex(v int) error {
	V := len(t.rank)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}
