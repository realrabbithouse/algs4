package graphapp

import (
	"algs4/graph"
	"fmt"
	"testing"
)

func TestTopological(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../testdata/tinyDAG.txt")
	if err != nil {
		t.Fatal(err)
	}
	top, _ := NewTopological(digraph)
	if top.HasOrder() {
		fmt.Println(top.Order())
	}
}
