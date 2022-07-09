package graphapp

import (
	"algs4/src/basic"
	"algs4/src/graph"
	"algs4/src/sort"
	"algs4/src/typ"
	"errors"
	"fmt"
	"math"
)

const PositiveInfinity = math.MaxFloat64

// ShortestPath is the interface that wrap the single-source shortest paths methods.
type ShortestPath interface {
	HashPathTo(id graph.ID) bool
	PathTo(id graph.ID) basic.Iterator
	DistTo(id graph.ID) float64
}

type NegativeWeightSP interface {
	ShortestPath
	HasNegativeCycle() bool
	NegativeCycle() basic.Iterator
}

// **************************************************************** //

// The DijkstraSP represents a data type for solving the single-source shortest paths
// problem in edge-weighted digraphs where the edge weights are non-negative.
type DijkstraSP struct {
	distTo []float64
	edgeTo []graph.DirectedEdge
	pq     *sort.IndexMinPQ
	G      *graph.EdgeWeightedDigraph
	source graph.ID
}

func NewDijkstraSP(G *graph.EdgeWeightedDigraph, s int) (*DijkstraSP, error) {
	if G == nil {
		return nil, errors.New("argument is nil")
	}
	V := G.V
	for i := 0; i < V; i++ {
		for _, e := range G.Adj(i) {
			if e.Weight() < 0 {
				return nil, errors.New(fmt.Sprintf("edge %v has negative weight", e))
			}
		}
	}
	if s < 0 || s >= V {
		return nil, errors.New(fmt.Sprintf("source vertex %d is not between 0 and %d", s, V))
	}

	distTo := make([]float64, V)
	edgeTo := make([]graph.DirectedEdge, V)
	for i := 0; i < V; i++ {
		distTo[i] = PositiveInfinity
	}
	distTo[s] = 0.0

	pq := sort.NewIndexMinPQWithCap(V)
	pq.Insert(s, typ.ComparableFloat64(distTo[s]))
	sp := DijkstraSP{
		distTo: distTo,
		edgeTo: edgeTo,
		pq:     pq,
		G:      G,
		source: graph.ID(s),
	}
	for !pq.IsEmpty() {
		v := pq.DelMin()
		for _, e := range G.Adj(v) {
			sp.relax(e)
		}
	}

	return &sp, nil
}

func (sp DijkstraSP) HasPathTo(id graph.ID) bool {
	if err := sp.validateIndex(int(id)); err != nil {
		fmt.Println("HasPathTo err:", err)
		return false
	}
	return sp.distTo[id] < PositiveInfinity
}

func (sp DijkstraSP) PathTo(id graph.ID) basic.Iterator {
	if err := sp.validateIndex(int(id)); err != nil {
		fmt.Println("PathTo err:", err)
		return nil
	}
	path := new(basic.Stack)
	var e graph.DirectedEdge
	for e = sp.edgeTo[id]; e.Src() != sp.source; e = sp.edgeTo[e.Src()] {
		tmp := e
		path.Push(tmp)
	}
	path.Push(e)
	return path
}

func (sp DijkstraSP) DistTo(id graph.ID) float64 {
	if err := sp.validateIndex(int(id)); err != nil {
		fmt.Println("DistTo err:", err)
		return 0.0
	}
	return sp.distTo[id]
}

func (sp DijkstraSP) validateIndex(v int) error {
	V := len(sp.distTo)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}

func (sp *DijkstraSP) relax(e graph.DirectedEdge) {
	v, w := e.Src(), e.Dst()
	if sp.distTo[w] > sp.distTo[v]+e.Weight() {
		sp.distTo[w] = sp.distTo[v] + e.Weight()
		sp.edgeTo[w] = e
		if sp.pq.Contains(int(w)) {
			sp.pq.Change(int(w), typ.ComparableFloat64(sp.distTo[w]))
		} else {
			sp.pq.Insert(int(w), typ.ComparableFloat64(sp.distTo[w]))
		}
	}
}

// **************************************************************** //

// The BellmanFordSP represents a data type for solving the single-source shortest
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
