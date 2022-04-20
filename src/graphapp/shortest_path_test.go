package graphapp

import (
	"algs4/src/graph"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestDijkstraSP(t *testing.T) {
	t0 := time.Now()
	digraph, err := graph.ReadEdgeWeightedDigraphFromFile("../data/largeEWD.txt")
	fmt.Println("graph loading time:", time.Since(t0))
	if err != nil {
		t.Fatal(err)
	}
	t1 := time.Now()
	sp, err := NewDijkstraSP(digraph, 0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("run time:", time.Since(t1))

	var n int
	for i := 0; i < 3; i++ {
		n = rand.Intn(100000)
		fmt.Printf("distTo[%d] = %.8f\n", n, sp.DistTo(graph.ID(n)))
		path := sp.PathTo(graph.ID(n))
		fmt.Println("================================================")
		for path.HasNext() {
			fmt.Println(path.Next())
		}
		fmt.Println("================================================")
	}
}
