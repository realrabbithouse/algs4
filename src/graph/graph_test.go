package graph

import (
	"fmt"
	"testing"
)

func TestGraph(t *testing.T) {
	g, err := ReadGraphFromFile("../data/tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(g)
	adj0 := g.Adj(0)
	adj1 := g.Adj(1)
	adj2 := g.Adj(2)
	deg0 := g.Degree(0)
	deg1 := g.Degree(1)
	deg2 := g.Degree(2)
	fmt.Println(adj0, deg0)
	fmt.Println(adj1, deg1)
	fmt.Println(adj2, deg2)
}
