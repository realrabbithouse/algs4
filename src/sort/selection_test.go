package sort

import (
	"fmt"
	"testing"
)

func TestSelection_Sort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	input := IntSlice{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	SelectionSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	fmt.Println(input)
}
