package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestDijkstraSP(t *testing.T) {
	digraph, err := graph.ReadEdgeWeightedDigraphFromFile("../data/tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	sp, err := NewDijkstraSP(digraph, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(sp.DistTo(1))
	fmt.Println(sp.DistTo(2))
	fmt.Println(sp.DistTo(3))
	path1 := sp.PathTo(1)
	fmt.Println("=====================")
	for path1.HasNext() {
		fmt.Println(path1.Next())
	}
	fmt.Println("=====================")
	path2 := sp.PathTo(2)
	for path2.HasNext() {
		fmt.Println(path2.Next())
	}
	fmt.Println("=====================")
	path3 := sp.PathTo(3)
	for path3.HasNext() {
		fmt.Println(path3.Next())
	}
}
