package graphapp

import (
	"algs4/src/basic"
)

// MST defines a minimum spanning tree that wraps the basic minimum
// spanning tree functionalities.
type MST interface {
	Weight() float64
	Edges() *basic.Queue
}

// **************************************************************** //
