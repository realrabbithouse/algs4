package graphapp

import (
	"errors"
	"fmt"

	"algs4/graph"
)

// DirectedDFS is mainly used for single- and multiple-source reachability.
type DirectedDFS struct {
	digraph *graph.Digraph
	marked  []bool
	count   int
}

func NewDirectedDFS(digraph *graph.Digraph, sources ...int) (*DirectedDFS, error) {
	if digraph == nil {
		return nil, errors.New("argument is nil")
	}
	d := DirectedDFS{
		digraph: digraph,
		marked:  make([]bool, digraph.V),
	}
	err := d.validateVertices(sources)
	if err != nil {
		return nil, err
	}
	for _, source := range sources {
		d.dfs(source)
	}
	return &d, nil
}

func (d *DirectedDFS) dfs(v int) {
	if !d.marked[v] {
		d.marked[v] = true
		d.count++
		for _, id := range d.digraph.Adj(v) {
			d.dfs(int(id))
		}
	}
}

func (d *DirectedDFS) Marked(v int) bool {
	if err := d.validateVertex(v); err != nil {
		fmt.Println("Marked err:", err)
		return false
	}
	return d.marked[v]
}

func (d *DirectedDFS) Count() int {
	return d.count
}

func (d *DirectedDFS) validateVertex(v int) error {
	V := len(d.marked)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}

func (d *DirectedDFS) validateVertices(vertices []int) error {
	if vertices == nil {
		return errors.New("argument is nil")
	}
	for _, id := range vertices {
		err := d.validateVertex(id)
		if err != nil {
			return err
		}
	}
	if len(vertices) == 0 {
		return errors.New("no vertices")
	}
	return nil
}
