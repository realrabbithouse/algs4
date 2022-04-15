package st

const (
	RED   = true
	BLACK = false
)

// colorNode in Red/Black binary search tree.
type colorNode struct {
	key   Key
	value Value
	left  *colorNode
	right *colorNode
	size  int
	color bool
}

type RedBlackBST struct {
	root *colorNode
}
