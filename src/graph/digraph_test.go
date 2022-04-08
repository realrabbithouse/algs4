package graph

import (
	"fmt"
	"testing"
)

func TestDigraph(t *testing.T) {
	digraph, err := ReadDigraphFromFile("../data/tinyG.txt")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(digraph)
	reverse := digraph.Reverse()
	fmt.Println(reverse)
	var tests = []struct {
		input  int
		wanted int
	}{
		{0, 4},
		{1, 0},
		{2, 0},
		{3, 0},
		{4, 1},
		{5, 2},
		{6, 1},
		{7, 1},
		{8, 0},
		{9, 3},
		{10, 0},
		{11, 1},
		{12, 0},
	}
	for _, test := range tests {
		if got := digraph.OutDegree(test.input); got != test.wanted {
			t.Errorf("OutDegree(%d) != %d", test.input, got)
		}
	}
}
