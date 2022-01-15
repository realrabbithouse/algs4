package sort

import (
	"algs4/src/prep"
	"fmt"
	"log"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	nums, err := prep.ReadNumFromFile("../prep/rand4096.out")
	input := IntSlice(nums)
	if err != nil {
		log.Fatal(err)
	}

	SelectionSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	fmt.Println(input)
}

func TestInsertionSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}
	nums, err := prep.ReadNumFromFile("../prep/rand4096.out")
	input := IntSlice(nums)
	if err != nil {
		log.Fatal(err)
	}

	InsertionSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	fmt.Println(input)
}

func TestBubbleSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}
	nums, err := prep.ReadNumFromFile("../prep/rand4096.out")
	input := IntSlice(nums)
	if err != nil {
		log.Fatal(err)
	}

	BubbleSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	fmt.Println(input)
}

func TestMergeSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := prep.ReadNumFromFile("../prep/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := IntSlice(prep.GenRandomNum(65536, 65536))

	MergeSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}
