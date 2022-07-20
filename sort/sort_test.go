package sort

import (
	"algs4/fileops"
	"algs4/typ"
	"fmt"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}

	//nums, err := fileops.ReadNumsFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(65536, 8192*8))

	SelectionSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}

func TestInsertionSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := fileops.ReadNumsFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(65536, 8192*8))

	InsertionSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}

func TestBubbleSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := fileops.ReadNumsFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(65536, 8192))

	BubbleSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}

func TestShellSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := fileops.ReadNumsFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(65536, 8192*8))

	ShellSort(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}

func TestMergeSort(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := fileops.ReadNumFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(65536, 65536))

	MergeSortInt(input)
	fmt.Println("is sorted?", input.IsSorted())
	//fmt.Println(input)
}

func TestQuick(t *testing.T) {
	//input := IntSlice{1, 32, 3, 52, 42, 563456, 34, 534, 98, 2, 234, 5345, 65, 2, 676}
	//input := IntSlice{10, 9, 9, 8, 7, 7, 6, 5, 5, 4, 3, 3, 2, 1}

	//nums, err := fileops.ReadNumFromFile("../fileops/rand4096.out")
	//input := IntSlice(nums)
	//if err != nil {
	//	log.Fatal(err)
	//}

	input := typ.IntSlice(fileops.GenRandomNums(1000000, 65536))

	QuickSortInt(input)
	fmt.Println("is sorted?", input.IsSorted())
	//for i := range input {
	//	fmt.Println(input[i])
	//}
}

func TestMaxPQ(t *testing.T) {
	input := typ.IntSlice(fileops.GenRandomNums(40000000, 20000000))
	//fmt.Println("input:", input)
	keys := make([]typ.Comparable, len(input))
	for i := range input {
		keys[i] = typ.ComparableInt(input[i])
	}
	maxPQ := NewMaxPQWithKeys(keys)
	fmt.Println("is binary heap ok?", maxPQ.isMaxHeap())
	fmt.Println("max element:", maxPQ.Max())
	fmt.Println("top 10:", maxPQ.TopK(10))
	fmt.Println("is still max heap?", maxPQ.isMaxHeap())
	fmt.Println("number of elements left:", maxPQ.Size())
}

func TestIndexMinPQ(t *testing.T) {
	input := typ.IntSlice(fileops.GenRandomNums(20, 5))
	//input = IntSlice{1, 2, 3, 4, 5}
	fmt.Println("input:", input)
	keys := make([]typ.Comparable, len(input))
	for i := range input {
		keys[i] = typ.ComparableInt(input[i])
	}
	minPQ := NewIndexMinPQWithKeys(keys)
	fmt.Println("qp:", minPQ.qp)
	fmt.Println("pq:", minPQ.pq)
	fmt.Println("keys:", minPQ.keys)
	fmt.Println("min key:", minPQ.MinKey())
	fmt.Println("min index:", minPQ.MinIndex())
	minPQ.Change(3, typ.ComparableInt(10010))
	fmt.Println("qp:", minPQ.qp)
	fmt.Println("pq:", minPQ.pq)
	fmt.Println("keys:", minPQ.keys)
	fmt.Println("min key:", minPQ.MinKey())
	fmt.Println("min index:", minPQ.MinIndex())
	//for i := 0; i < len(keys); i++ {
	//	fmt.Println(i, minPQ.KeyOf(i))
	//}
	fmt.Println("delete min:", minPQ.DelMin())
	fmt.Println("qp:", minPQ.qp)
	fmt.Println("pq:", minPQ.pq)
	fmt.Println("keys:", minPQ.keys)
	fmt.Println("delete min:", minPQ.DelMin())
	fmt.Println("qp:", minPQ.qp)
	fmt.Println("pq:", minPQ.pq)
	fmt.Println("keys:", minPQ.keys)
}

func TestHeap(t *testing.T) {
	input := typ.IntSlice(fileops.GenRandomNums(4000000, 2000000))
	//fmt.Println("input:", input)
	keys := make([]typ.Comparable, len(input))
	for i := range input {
		keys[i] = typ.ComparableInt(input[i])
	}
	h := Heap{pq: keys}
	h.Sort()
	h.Show()
}

func TestMinPQ(t *testing.T) {
	input := typ.IntSlice(fileops.GenRandomNums(100, 50))
	//fmt.Println("input:", input)
	keys := make([]typ.Comparable, len(input))
	for i := range input {
		keys[i] = typ.ComparableInt(input[i])
	}
	minPQ := NewMinPQWithKeys(keys)

	//for !minPQ.IsEmpty() {
	//	fmt.Println(minPQ.DelMin())
	//}

	fmt.Println("check result:", minPQ.isMinHeap())
	fmt.Println("min:", minPQ.Min())
	fmt.Println("min 49:", minPQ.MinK(49))
	fmt.Println("check result after deleting 20:", minPQ.isMinHeap())
}
