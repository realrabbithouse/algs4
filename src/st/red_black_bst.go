package st

const (
	RED   = true
	BLACK = false
)

// ColorNode in Red/Black binary search tree.
type ColorNode struct {
	key   Key
	value Value
	left  *ColorNode
	right *ColorNode
	sz    int
	color bool
}

type RedBlackBST struct {
	root *ColorNode
}
