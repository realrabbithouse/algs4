package graphapp

import (
	"errors"

	"algs4/basic"
	"algs4/graph"
)

type DirectedCycle struct {
	digraph *graph.Digraph
	edgeTo  []graph.ID
	marked  []bool
	onStack []bool
	cycle   *basic.Stack
}

func NewDirectedCycle(digraph *graph.Digraph) (*DirectedCycle, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	cycle := DirectedCycle{
		digraph: digraph,
		edgeTo:  make([]graph.ID, digraph.V),
		marked:  make([]bool, digraph.V),
		onStack: make([]bool, digraph.V),
	}
	for i := 0; i < digraph.V; i++ {
		if !cycle.marked[i] && cycle.cycle == nil {
			cycle.dfs(i)
		}
	}
	return &cycle, nil
}

func (c *DirectedCycle) HasCycle() bool {
	return c.cycle != nil
}

func (c *DirectedCycle) Cycle() []graph.ID {
	cycle := make([]graph.ID, 0, c.cycle.Size())
	var iter basic.Iterator = c.cycle
	for iter.HasNext() {
		v := iter.Next()
		cycle = append(cycle, v.(graph.ID))
	}
	return cycle
}

// dfs runs DFS and finds a directed cycle (if one exists)
func (c *DirectedCycle) dfs(v int) {
	if !c.marked[v] {
		c.marked[v] = true
		c.onStack[v] = true
		for _, w := range c.digraph.Adj(v) {
			// Short circuit if directed cycle found.
			if c.cycle != nil {
				return
			}
			if !c.marked[w] {
				c.edgeTo[w] = graph.ID(v)
				c.dfs(int(w))
			} else if c.onStack[w] { // A node cannot be on stack if it is not marked yet.
				c.cycle = new(basic.Stack)
				for i := graph.ID(v); i != w; i = c.edgeTo[i] {
					c.cycle.Push(i)
				}
				c.cycle.Push(w)
				c.cycle.Push(graph.ID(v))
			}
		}
		c.onStack[v] = false
	}
}
