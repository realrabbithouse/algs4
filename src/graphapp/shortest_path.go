package graphapp

import (
	"algs4/src/basic"
	"algs4/src/graph"
	"math"
)

const PositiveInfinity = math.MaxFloat64

// ShortestPath is the interface that wrap the single-source shortest paths methods.
type ShortestPath interface {
	HashPathTo(id graph.ID) bool
	PathTo(id graph.ID) []graph.Edge
	DistTo(id graph.ID) float64
}

type NegativeWeightSP interface {
	ShortestPath
	HasNegativeCycle() bool
	NegativeCycle() []graph.Edge
}

// **************************************************************** //

// The BellmanFordSP class represents a data type for solving the single-source shortest
// paths problem in edge-weighted digraphs with no negative cycles. The edge weights can
// be positive, negative, or zero. This class finds either a shortest path from the source
// vertex s to every other vertex or a negative cycle reachable from the source vertex.
//
// This correctly computes shortest paths if all arithmetic performed is without
// floating-point rounding error or arithmetic overflow. This is the case if all
// edge weights are integers and if none of the intermediate results exceeds 2^52.
type BellmanFordSP struct {
	distTo  []float64            // distTo[v] = distance  of shortest s->v path
	edgeTo  []graph.DirectedEdge // edgeTo[v] = last edge on shortest s->v path
	onQueue []bool               // onQueue[v] = is v currently on the queue?
	queue   *basic.Queue         // queue of vertices to relax
	cost    int                  // number of calls to relax()
	cycle   *basic.Stack         // negative cycle (or null if no such cycle)
}
