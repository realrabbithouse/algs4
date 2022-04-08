package graph

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedGraph(t *testing.T) {
	graph, err := ReadEdgeWeightedGraphFromFile("../data/tinyEWG.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(graph)
}
