package str

import (
	"fmt"
	"testing"
)

func TestBoyerMoore(t *testing.T) {
	bm := NewBoyerMoore("rabbit")
	n := bm.Search("when rrr rabbit rules the world, a lot of rrr rabbits...")
	fmt.Println("target hit at", n)
}
