package basic

import (
	"fmt"
	"testing"
)

func TestPrimeFactor(t *testing.T) {
	input := []uint64{10, 20, 30, 40, 50, 13, 33, 53, 73, 93}
	for i := range input {
		fmt.Println(PrimeFactor(input[i]))
	}
}
