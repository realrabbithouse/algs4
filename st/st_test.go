package st

import (
	"algs4/typ"
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}
	bst.Put(typ.ComparableString("hello"), "Hello")
	bst.Put(typ.ComparableString("world"), "World")
	bst.Put(typ.ComparableString("apple"), "Apple")
	bst.Put(typ.ComparableString("banana"), "Banana")
	bst.Put(typ.ComparableString("cat"), "Cat")
	bst.Put(typ.ComparableString("dodge"), "Dodge")
	bst.Put(typ.ComparableString("earth"), "Earth")
	bst.Put(typ.ComparableString("father"), "Father")
	bst.Put(typ.ComparableString("great"), "Great")
	bst.Put(typ.ComparableString("help"), "Help")
	bst.Put(typ.ComparableString("i"), "I")
	bst.Put(typ.ComparableString("job"), "Job")
	bst.Put(typ.ComparableString("kick"), "Kick")
	fmt.Println("min:", bst.Min())
	fmt.Println("max:", bst.Max())
	fmt.Println(bst.ToSlice())
	bst.Delete(typ.ComparableString("hello"))
	bst.Delete(typ.ComparableString("world"))
	fmt.Println(bst.ToSlice())
	fmt.Println("floor of danger:", bst.Floor(typ.ComparableString("danger")))
	fmt.Println("ceiling of danger:", bst.Ceiling(typ.ComparableString("danger")))
	fmt.Println("rank of danger:", bst.Rank(typ.ComparableString("danger")))
	fmt.Println("choose 3:", bst.Choose(3))
	fmt.Println("kick:", bst.Get(typ.ComparableString("kick")))
	fmt.Println("======= fileops order =======")
	bst.PreOrder()
	fmt.Println("======= in order =======")
	bst.InOrder()
	fmt.Println("======= post order =======")
	bst.PostOrder()
}
