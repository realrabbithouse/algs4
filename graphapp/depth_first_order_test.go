package graphapp

import (
	"algs4/graph"
	"fmt"
	"testing"
)

func TestDepthFirstOrder(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../testdata/tinyDAG.txt")
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
