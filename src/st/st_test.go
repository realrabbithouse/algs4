package st

import (
	"algs4/src/rabbit"
	"fmt"
	"testing"
)

func TestBST(t *testing.T) {
	bst := BST{}
	bst.Put(rabbit.ComparableString("hello"), "Hello")
	bst.Put(rabbit.ComparableString("world"), "World")
	bst.Put(rabbit.ComparableString("apple"), "Apple")
	bst.Put(rabbit.ComparableString("banana"), "Banana")
	bst.Put(rabbit.ComparableString("cat"), "Cat")
	bst.Put(rabbit.ComparableString("dodge"), "Dodge")
	bst.Put(rabbit.ComparableString("earth"), "Earth")
	bst.Put(rabbit.ComparableString("father"), "Father")
	bst.Put(rabbit.ComparableString("great"), "Great")
	bst.Put(rabbit.ComparableString("help"), "Help")
	bst.Put(rabbit.ComparableString("i"), "I")
	bst.Put(rabbit.ComparableString("job"), "Job")
	bst.Put(rabbit.ComparableString("kick"), "Kick")
	fmt.Println("min:", bst.Min())
	fmt.Println("max:", bst.Max())
	fmt.Println(bst.ToSlice())
	bst.Delete(rabbit.ComparableString("hello"))
	bst.Delete(rabbit.ComparableString("world"))
	fmt.Println(bst.ToSlice())
	fmt.Println("floor of danger:", bst.Floor(rabbit.ComparableString("danger")))
	fmt.Println("ceiling of danger:", bst.Ceiling(rabbit.ComparableString("danger")))
	fmt.Println("rank of danger:", bst.Rank(rabbit.ComparableString("danger")))
	fmt.Println("choose 3:", bst.Choose(3))
	fmt.Println("kick:", bst.Get(rabbit.ComparableString("kick")))
	fmt.Println("======= pre order =======")
	bst.PreOrder()
	fmt.Println("======= in order =======")
	bst.InOrder()
	fmt.Println("======= post order =======")
	bst.PostOrder()
}
