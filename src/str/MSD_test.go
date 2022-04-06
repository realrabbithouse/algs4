package str

import (
	"fmt"
	"testing"
)

func TestMSDSort(t *testing.T) {
	a := GenRandomStrings(100, 20)
	MSDSort(a)
	fmt.Println(a)
}
