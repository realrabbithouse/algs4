package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestDirectedDFS(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../data/tinyDG.txt")
	if err != nil {
		t.Fatal(err)
	}
	for i := 0; i < digraph.V; i++ {
		d, err := NewDirectedDFS(digraph, i)
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(d.Count())
	}
}
