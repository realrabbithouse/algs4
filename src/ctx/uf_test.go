package ctx

import (
	"fmt"
	"testing"
)

func TestUnionFind(t *testing.T) {
	uf, _ := NewUnionFind(15)
	uf.Union(0, 1)
	uf.Union(0, 3)
	uf.Union(0, 5)
	uf.Union(0, 7)
	uf.Union(0, 9)
	uf.Union(2, 4)
	uf.Union(12, 14)
	uf.Union(4, 6)
	uf.Union(10, 12)
	fmt.Println(uf.Count())
	fmt.Println(uf.Connected(1, 3))
	fmt.Println(uf.Connected(1, 5))
	fmt.Println(uf.Connected(1, 7))
	fmt.Println(uf.Connected(2, 3))
	fmt.Println(uf.Connected(2, 12))
	fmt.Println(uf.Connected(2, 6))
}
