package basic

import (
	"fmt"
	"testing"
)

func TestBag(t *testing.T) {
	bag := Bag{first: nil, n: 0}
	bag.Add("hello")
	bag.Add("world")
	bag.Add("rabbit")
	bag.Add("panda")
	bag.Add("tiger")
	for bag.HasItem() {
		fmt.Printf("%s ", bag.Next())
	}
	fmt.Println("\nis bag empty?", bag.IsEmpty())
	fmt.Println("bag size:", bag.Size())

	ibag := Bag{}
	ibag.Add(1)
	ibag.Add(2)
	ibag.Add(3)
	ibag.Add(4)
	ibag.Add(5)
	for ibag.HasItem() {
		fmt.Printf("%d ", ibag.Next())
	}
}
