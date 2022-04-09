package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"testing"
)

func TestSCC(t *testing.T) {
	digraph, err := graph.ReadDigraphFromFile("../data/tinyDG.txt")
	if err != nil {
		t.Fatal(err)
	}
	scc, _ := NewKosarajuSharirSCC(digraph)
	fmt.Println("Number of strongly connected components:", scc.count)
	fmt.Println("SCC information:", scc.id)
}
