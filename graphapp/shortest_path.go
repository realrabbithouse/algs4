package graphapp

import (
	basic2 "algs4/basic"
	graph2 "algs4/graph"
	"algs4/sort"
	"algs4/typ"
	"errors"
	"fmt"
	"math"
)

const PositiveInfinity = math.MaxFloat64

// ShortestPath is the interface that wrap the single-source shortest paths methods.
type ShortestPath interface {
	HashPathTo(id graph2.ID) bool
	PathTo(id graph2.ID) basic2.Iterator
	DistTo(id graph2.ID) float64
}

type NegativeWeightSP interface {
	ShortestPath
	HasNegativeCycle() bool
	NegativeCycle() basic2.Iterator
}

// **************************************************************** //

// The DijkstraSP represents a testdata type for solving the single-source shortest paths
// problem in edge-weighted digraphs where the edge weights are non-negative.
type DijkstraSP struct {
	distTo []float64
	edgeTo []graph2.DirectedEdge
	pq     *sort.IndexMinPQ
	G      *graph2.EdgeWeightedDigraph
	source graph2.ID
}

func NewDijkstraSP(G *graph2.EdgeWeightedDigraph, s int) (*DijkstraSP, error) {
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
	edgeTo := make([]graph2.DirectedEdge, V)
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
		source: graph2.ID(s),
	}
	for !pq.IsEmpty() {
		v := pq.DelMin()
		for _, e := range G.Adj(v) {
			sp.relax(e)
		}
	}

	return &sp, nil
}

func (sp DijkstraSP) HasPathTo(id graph2.ID) bool {
	if err := sp.validateIndex(int(id)); err != nil {
		fmt.Println("HasPathTo err:", err)
		return false
	}
	return sp.distTo[id] < PositiveInfinity
}

func (sp DijkstraSP) PathTo(id graph2.ID) basic2.Iterator {
	if err := sp.validateIndex(int(id)); err != nil {
		fmt.Println("PathTo err:", err)
		return nil
	}
	path := new(basic2.Stack)
	var e graph2.DirectedEdge
	for e = sp.edgeTo[id]; e.Src() != sp.source; e = sp.edgeTo[e.Src()] {
		tmp := e
		path.Push(tmp)
	}
	path.Push(e)
	return path
}

func (sp DijkstraSP) DistTo(id graph2.ID) float64 {
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

func (sp *DijkstraSP) relax(e graph2.DirectedEdge) {
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

// The BellmanFordSP represents a testdata type for solving the single-source shortest
// paths problem in edge-weighted digraphs with no negative cycles. The edge weights can
// be positive, negative, or zero. This class finds either a shortest path from the source
// vertex s to every other vertex or a negative cycle reachable from the source vertex.
//
// This correctly computes shortest paths if all arithmetic performed is without
// floating-point rounding error or arithmetic overflow. This is the case if all
// edge weights are integers and if none of the intermediate results exceeds 2^52.
type BellmanFordSP struct {
	distTo  []float64             // distTo[v] = distance  of shortest s->v path
	edgeTo  []graph2.DirectedEdge // edgeTo[v] = last edge on shortest s->v path
	onQueue []bool                // onQueue[v] = is v currently on the queue?
	queue   *basic2.Queue         // queue of vertices to relax
	cost    int                   // number of calls to relax()
	cycle   *basic2.Stack         // negative cycle (or null if no such cycle)
}
