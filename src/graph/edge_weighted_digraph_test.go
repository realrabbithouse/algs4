package graph

import (
	"fmt"
	"testing"
)

func TestEdgeWeightedDigraph(t *testing.T) {
	graph, err := ReadEdgeWeightedDigraphFromFile("../data/tinyEWD.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(graph)
	reverse := graph.Reverse()
	fmt.Println(reverse)
	var tests = []int{2, 1, 1, 1, 2, 3, 3, 2}
	for i, wanted := range tests {
		fmt.Printf("adj %d: %v\n", i, graph.Adj(i))
		if got := graph.OutDegree(i); got != wanted {
			t.Errorf("OutDegree(%d) != %d", got, wanted)
		}
	}
}
