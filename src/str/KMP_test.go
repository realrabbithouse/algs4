package str

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	kmp := NewKMP("AACAA")
	n := kmp.Search("AABRAACADABRAACAADABRA")
	fmt.Println("target hit at", n)

	n = kmp.Search("when rrr typ rules the world, a lot of rrr rabbits...")
	fmt.Println("target hit at", n)
}
