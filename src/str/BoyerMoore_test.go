package str

import (
	"fmt"
	"testing"
)

func TestBoyerMoore(t *testing.T) {
	bm := NewBoyerMoore("typ")
	n := bm.Search("when rrr typ rules the world, a lot of rrr rabbits...")
	fmt.Println("target hit at", n)
}
