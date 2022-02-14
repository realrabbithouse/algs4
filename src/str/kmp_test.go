package str

import (
	"fmt"
	"testing"
)

func TestKMP(t *testing.T) {
	kmp := NewKMP("AACAA")
	n := kmp.Search("AABRAACADABRAACAADABRA")
	fmt.Println("target hit at", n)
}
