package basic

type Stack struct {
	head *linked
	iter *linked
	n    int
}

func NewStack() *Stack {
	return new(Stack)
}

func (stack *Stack) Push(val interface{}) {
	newHead := &linked{
		item: val,
		next: stack.head,
	}
	stack.head = newHead
	stack.iter = stack.head
	stack.n++
}

func (stack *Stack) Pop() interface{} {
	if stack.head == nil {
		return nil
	}
	val := stack.head.item
	stack.head = stack.head.next
	stack.iter = stack.head
	stack.n--
	return val
}

func (stack *Stack) IsEmpty() bool {
	return stack.head == nil
}

func (stack *Stack) Size() int {
	return stack.n
}

func (stack *Stack) HasNext() bool {
	if stack.iter != nil {
		return true
	}
	stack.iter = stack.head
	return false
}

func (stack *Stack) Next() interface{} {
	val := stack.iter.item
	stack.iter = stack.iter.next
	return val
}
