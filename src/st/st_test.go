package st

import (
	"algs4/src/sort"
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}
	bst.Put(sort.ComparableString("hello"), "Hello")
	bst.Put(sort.ComparableString("world"), "World")
	bst.Put(sort.ComparableString("apple"), "Apple")
	bst.Put(sort.ComparableString("banana"), "Banana")
	bst.Put(sort.ComparableString("cat"), "Cat")
	bst.Put(sort.ComparableString("dodge"), "Dodge")
	bst.Put(sort.ComparableString("earth"), "Earth")
	bst.Put(sort.ComparableString("father"), "Father")
	bst.Put(sort.ComparableString("great"), "Great")
	bst.Put(sort.ComparableString("help"), "Help")
	bst.Put(sort.ComparableString("i"), "I")
	bst.Put(sort.ComparableString("job"), "Job")
	bst.Put(sort.ComparableString("kick"), "Kick")
	fmt.Println("min:", bst.Min())
	fmt.Println("max:", bst.Max())
	fmt.Println(bst.ToSlice())
	bst.Delete(sort.ComparableString("hello"))
	bst.Delete(sort.ComparableString("world"))
	fmt.Println(bst.ToSlice())
	fmt.Println("floor of danger:", bst.Floor(sort.ComparableString("danger")))
	fmt.Println("ceiling of danger:", bst.Ceiling(sort.ComparableString("danger")))
	fmt.Println("rank of danger:", bst.Rank(sort.ComparableString("danger")))
	fmt.Println("choose 3:", bst.Choose(3))
	fmt.Println("kick:", bst.Get(sort.ComparableString("kick")))
}
