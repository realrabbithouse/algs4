package basic

import (
	"fmt"
	"testing"
)

func TestBag(t *testing.T) {
	var it Iterator
	bag := Bag{first: nil, n: 0}
	bag.Add("hello")
	bag.Add("world")
	bag.Add("typ")
	bag.Add("panda")
	bag.Add("tiger")
	fmt.Println("is bag empty?", bag.IsEmpty())
	fmt.Println("bag size:", bag.Size())
	it = &bag
	for it.HasNext() {
		fmt.Printf("%s ", it.Next())
	}
	fmt.Print("\n")

	ibag := Bag{}
	ibag.Add(1)
	ibag.Add(2)
	ibag.Add(3)
	ibag.Add(4)
	ibag.Add(5)
	it = &ibag
	for it.HasNext() {
		fmt.Printf("%d ", it.Next())
	}
	fmt.Print("\n")
}
