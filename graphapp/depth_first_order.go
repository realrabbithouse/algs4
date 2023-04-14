package graphapp

import (
	"errors"
	"fmt"

	"algs4/basic"
	"algs4/graph"
)

type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    *basic.Queue
	postorder   *basic.Queue
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(digraph *graph.Digraph) (*DepthFirstOrder, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	dfo := DepthFirstOrder{
		marked:    make([]bool, digraph.V),
		pre:       make([]int, digraph.V),
		post:      make([]int, digraph.V),
		preorder:  new(basic.Queue),
		postorder: new(basic.Queue),
	}
	for i := 0; i < digraph.V; i++ {
		if !dfo.marked[i] {
			dfo.dfs(digraph, i)
		}
	}
	return &dfo, nil
}

func (dfo *DepthFirstOrder) Preorder() []graph.ID {
	preorder := make([]graph.ID, 0, dfo.preorder.Size())
	var iter basic.Iterator = dfo.preorder
	for iter.HasNext() {
		preorder = append(preorder, iter.Next().(graph.ID))
	}
	return preorder
}

func (dfo *DepthFirstOrder) Postorder() []graph.ID {
	postorder := make([]graph.ID, 0, dfo.postorder.Size())
	var iter basic.Iterator = dfo.postorder
	for iter.HasNext() {
		postorder = append(postorder, iter.Next().(graph.ID))
	}
	return postorder
}

func (dfo *DepthFirstOrder) ReversePostOrder() []graph.ID {
	reverse := make([]graph.ID, 0, dfo.postorder.Size())
	stack := new(basic.Stack)
	var iter basic.Iterator = dfo.postorder
	for iter.HasNext() {
		stack.Push(iter.Next())
	}
	iter = stack
	for iter.HasNext() {
		reverse = append(reverse, iter.Next().(graph.ID))
	}
	return reverse
}

func (dfo *DepthFirstOrder) Pre(v int) int {
	err := dfo.validateVertex(v)
	if err != nil {
		fmt.Println("Pre err:", err)
		return -1
	}
	return dfo.pre[v]
}

func (dfo *DepthFirstOrder) Post(v int) int {
	err := dfo.validateVertex(v)
	if err != nil {
		fmt.Println("Post err:", err)
		return -1
	}
	return dfo.post[v]
}

func (dfo *DepthFirstOrder) dfs(digraph *graph.Digraph, v int) {
	dfo.marked[v] = true
	dfo.preorder.Enqueue(graph.ID(v))
	dfo.pre[v] = dfo.preCounter
	dfo.preCounter++
	for _, w := range digraph.Adj(v) {
		if !dfo.marked[w] {
			dfo.dfs(digraph, int(w))
		}
	}
	dfo.postorder.Enqueue(graph.ID(v))
	dfo.post[v] = dfo.postCounter
	dfo.postCounter++
}

func (dfo *DepthFirstOrder) validateVertex(v int) error {
	V := len(dfo.marked)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}
