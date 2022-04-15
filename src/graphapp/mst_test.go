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

func TestMST(t *testing.T) {
	g, err := graph.ReadEdgeWeightedGraphFromFile("../data/tinyEWG.txt")
	if err != nil {
		t.Fatal(err)
	}

	var mst MST
	mst, _ = NewKruskalMST(g)
	fmt.Println(mst.Edges())
	fmt.Println(mst.Weight())

	mst, _ = NewLazyPrimMST(g)
	fmt.Println(mst.Edges())
	fmt.Println(mst.Weight())
}
