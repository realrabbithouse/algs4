package graphapp

import (
	basic2 "algs4/basic"
	graph2 "algs4/graph"
	"errors"
	"fmt"
)

type DepthFirstOrder struct {
	marked      []bool
	pre         []int
	post        []int
	preorder    *basic2.Queue
	postorder   *basic2.Queue
	preCounter  int
	postCounter int
}

func NewDepthFirstOrder(digraph *graph2.Digraph) (*DepthFirstOrder, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	dfo := DepthFirstOrder{
		marked:    make([]bool, digraph.V),
		pre:       make([]int, digraph.V),
		post:      make([]int, digraph.V),
		preorder:  new(basic2.Queue),
		postorder: new(basic2.Queue),
	}
	for i := 0; i < digraph.V; i++ {
		if !dfo.marked[i] {
			dfo.dfs(digraph, i)
		}
	}
	return &dfo, nil
}

func (dfo *DepthFirstOrder) Preorder() []graph2.ID {
	preorder := make([]graph2.ID, 0, dfo.preorder.Size())
	var iter basic2.Iterator = dfo.preorder
	for iter.HasNext() {
		preorder = append(preorder, iter.Next().(graph2.ID))
	}
	return preorder
}

func (dfo *DepthFirstOrder) Postorder() []graph2.ID {
	postorder := make([]graph2.ID, 0, dfo.postorder.Size())
	var iter basic2.Iterator = dfo.postorder
	for iter.HasNext() {
		postorder = append(postorder, iter.Next().(graph2.ID))
	}
	return postorder
}

func (dfo *DepthFirstOrder) ReversePostOrder() []graph2.ID {
	reverse := make([]graph2.ID, 0, dfo.postorder.Size())
	stack := new(basic2.Stack)
	var iter basic2.Iterator = dfo.postorder
	for iter.HasNext() {
		stack.Push(iter.Next())
	}
	iter = stack
	for iter.HasNext() {
		reverse = append(reverse, iter.Next().(graph2.ID))
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

func (dfo *DepthFirstOrder) dfs(digraph *graph2.Digraph, v int) {
	dfo.marked[v] = true
	dfo.preorder.Enqueue(graph2.ID(v))
	dfo.pre[v] = dfo.preCounter
	dfo.preCounter++
	for _, w := range digraph.Adj(v) {
		if !dfo.marked[w] {
			dfo.dfs(digraph, int(w))
		}
	}
	dfo.postorder.Enqueue(graph2.ID(v))
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
