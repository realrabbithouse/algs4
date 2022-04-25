package classic

import (
	"fmt"
	"testing"
)

func TestEightQueen(t *testing.T) {
	res := EightQueen()
	for res.HasNext() {
		fmt.Println(res.Next())
	}
}

func TestPermutation(t *testing.T) {
	a := []int{1, 2, 3, 4}
	Permutation(a)
}
