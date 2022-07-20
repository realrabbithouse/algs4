package graphapp

import (
	graph2 "algs4/graph"
	"errors"
	"fmt"
)

//
// Kahn’s algorithm
//
// L <- Empty list that will contain the sorted elements
// S <- Set of all nodes with no incoming edge
//
// while S is not empty do
//    remove a node N from S
//    add N to L
//    for each node M with an edge E from N to M do
//        remove edge E from the graph
//        if M has no other incoming edges then
//            insert M into S
//
// if graph has edges then
//    return error  # graph has at least one cycle
// else
//    return L  # a topologically sorted order
//

// Topological 拓扑排序
type Topological struct {
	order []graph2.ID
	rank  []int
}

func NewTopological(digraph *graph2.Digraph) (*Topological, error) {
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

func (t *Topological) Order() []graph2.ID {
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
