package graphapp

import (
	"algs4/graph"
	"fmt"
	"testing"
)

func TestDirectedCycle(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../testdata/tinyDG.txt")
	if err != nil {
		t.Fatal(err)
	}
	cycle, err := NewDirectedCycle(digraph)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cycle.Cycle())
}
