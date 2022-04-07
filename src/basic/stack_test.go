package basic

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	var iter Iterator = stack
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}

	fmt.Println("stack size:", stack.Size())
	fmt.Println("is empty:", stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println("stack size:", stack.Size())
	fmt.Println("is empty:", stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println("stack size:", stack.Size())
	fmt.Println("is empty:", stack.IsEmpty())
}
