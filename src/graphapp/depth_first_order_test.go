package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestDepthFirstOrder(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../data/tinyDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	dfo, err := NewDepthFirstOrder(digraph)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(dfo.Preorder())
	fmt.Println(dfo.Postorder())
	fmt.Println(dfo.ReversePostOrder())
}
