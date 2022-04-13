package basic

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := new(Queue)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)

	var iter Iterator = queue
	for iter.HasNext() {
		fmt.Println(iter.Next())
	}

	fmt.Println("size:", queue.Size())
	fmt.Println("is empty:", queue.IsEmpty())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println("size:", queue.Size())
	fmt.Println("is empty:", queue.IsEmpty())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println("size:", queue.Size())
	fmt.Println("is empty:", queue.IsEmpty())

	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println(queue.Dequeue())
	fmt.Println("size:", queue.Size())
	fmt.Println("is empty:", queue.IsEmpty())
}
