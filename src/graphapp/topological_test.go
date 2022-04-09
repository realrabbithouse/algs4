package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestTopological(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../data/tinyDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	top, _ := NewTopological(digraph)
	if top.HasOrder() {
		fmt.Println(top.Order())
	}
}
