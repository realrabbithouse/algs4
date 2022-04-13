package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestLazyPrimMST(t *testing.T) {
	g, err := graph.ReadEdgeWeightedGraphFromFile("../data/tinyEWG.txt")
	if err != nil {
		t.Fatal(err)
	}
	prim, err := NewLazyPrimMST(g)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(prim.Edges())
	fmt.Println(prim.Weight())
}
