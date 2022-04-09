package graphapp

import (
	"algs4/src/graph"
	"errors"
	"fmt"
)

type SCC interface {
	StronglyConnected(v, w graph.ID) bool
	Count() int
	Id(v graph.ID) int
}

// KosarajuSharirSCC algorithm: run DFS on G, using reverse postorder to guide calculation.
type KosarajuSharirSCC struct {
	marked []bool
	id     []int
	count  int
}

func NewKosarajuSharirSCC(digraph *graph.Digraph) (*KosarajuSharirSCC, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	dfo, _ := NewDepthFirstOrder(digraph.Reverse())
	ssc := KosarajuSharirSCC{
		marked: make([]bool, digraph.V),
		id:     make([]int, digraph.V),
	}
	for _, v := range dfo.ReversePostOrder() {
		if !ssc.marked[v] {
			ssc.dfs(digraph, int(v))
			ssc.count++
		}
	}
	return &ssc, nil
}

func (scc KosarajuSharirSCC) StronglyConnected(v, w graph.ID) bool {
	var err error
	err = scc.validateVertex(int(v))
	if err != nil {
		fmt.Println("StronglyConnected err:", err)
		return false
	}
	err = scc.validateVertex(int(w))
	if err != nil {
		fmt.Println("StronglyConnected err:", err)
		return false
	}
	return scc.id[v] == scc.id[w]
}

func (scc KosarajuSharirSCC) Count() int {
	return scc.count
}

func (scc KosarajuSharirSCC) Id(v graph.ID) int {
	err := scc.validateVertex(int(v))
	if err != nil {
		fmt.Println("Id err:", err)
		return -1
	}
	return scc.id[v]
}

func (scc *KosarajuSharirSCC) dfs(digraph *graph.Digraph, v int) {
	scc.marked[v] = true
	scc.id[v] = scc.count
	for _, w := range digraph.Adj(v) {
		if !scc.marked[w] {
			scc.dfs(digraph, int(w))
		}
	}
}

func (scc *KosarajuSharirSCC) validateVertex(v int) error {
	V := len(scc.marked)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}
