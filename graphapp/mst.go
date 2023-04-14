package graphapp

import (
	"errors"
	"fmt"
	ofsort "sort"

	"algs4/basic"
	"algs4/graph"
	"algs4/sort"
	"algs4/uf"
)

// MST defines a minimum spanning tree that wraps the basic minimum
// spanning tree functionalities.
type MST interface {
	Weight() float64
	Edges() []graph.Edge
}

// **************************************************************** //

type LazyPrimMST struct {
	marked []bool       // marked[v] = true iff v on tree
	mst    *basic.Queue // edges in the MST
	pq     *sort.MinPQ  // edges with one endpoint in tree
	weight float64      // total weight of MST
}

func NewLazyPrimMST(G *graph.EdgeWeightedGraph) (*LazyPrimMST, error) {
	if G == nil {
		return nil, errors.New("argument is nil")
	}
	prim := LazyPrimMST{
		marked: make([]bool, G.V),
		mst:    new(basic.Queue),
		pq:     sort.NewMinPQWithCap(G.V),
	}
	// Run Prim from all vertices to get a minimum spanning forest.
	for i := 0; i < G.V; i++ {
		if !prim.marked[i] {
			prim.visit(G, i)
			for !prim.pq.IsEmpty() {
				edge := prim.pq.DelMin().(graph.Edge)
				v := edge.Either()
				w, _ := edge.Other(v)
				if prim.marked[v] && prim.marked[w] {
					continue
				}
				prim.mst.Enqueue(edge)
				prim.weight += edge.Weight()
				if prim.marked[v] {
					prim.visit(G, int(w))
				} else {
					prim.visit(G, int(v))
				}
			}
		}
	}
	return &prim, nil
}

func (p LazyPrimMST) Edges() []graph.Edge {
	N := p.mst.Size()
	var iter basic.Iterator = p.mst
	edges := make([]graph.Edge, 0, N)
	for iter.HasNext() {
		edges = append(edges, iter.Next().(graph.Edge))
	}
	return edges
}

func (p LazyPrimMST) Weight() float64 {
	return p.weight
}

func (p *LazyPrimMST) visit(G *graph.EdgeWeightedGraph, v int) {
	if err := p.validateIndex(v); err != nil {
		fmt.Println("Invalid index:", err)
		return
	}
	p.marked[v] = true
	for _, edge := range G.Adj(v) {
		w, _ := edge.Other(graph.ID(v))
		if !p.marked[w] {
			p.pq.Insert(edge)
		}
	}
}

func (p *LazyPrimMST) validateIndex(v int) error {
	V := len(p.marked)
	if v < 0 || v >= V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", v, V))
	}
	return nil
}

// The KruskalMST represents a testdata type for computing a minimum spanning tree in an edge-weighted graph.
// The edge weights can be positive, zero, or negative and need not be distinct. If the graph is not
// connected, it computes a minimum spanning forest, which is the union of minimum spanning trees in
// each connected component.
type KruskalMST struct {
	mst    *basic.Queue
	weight float64
}

func NewKruskalMST(G *graph.EdgeWeightedGraph) (*KruskalMST, error) {
	if G == nil {
		return nil, errors.New("argument is nil")
	}
	uf, err := uf.NewUnionFind(G.V)
	if err != nil {
		return nil, err
	}
	kruskal := KruskalMST{
		mst: new(basic.Queue),
	}
	edges := make([]graph.Edge, G.E)
	copy(edges, G.Edges())
	ofsort.Slice(edges, func(i, j int) bool {
		return edges[i].CompareTo(edges[j]) < 0
	})
	for i := range edges {
		edge := edges[i]
		v := edge.Either()
		w, _ := edge.Other(v)
		if !uf.Connected(int(v), int(w)) {
			kruskal.mst.Enqueue(edge)
			kruskal.weight += edge.Weight()
			uf.Union(int(v), int(w))
		}
	}
	return &kruskal, nil
}

func (k *KruskalMST) Edges() []graph.Edge {
	N := k.mst.Size()
	var iter basic.Iterator = k.mst
	edges := make([]graph.Edge, 0, N)
	for iter.HasNext() {
		edges = append(edges, iter.Next().(graph.Edge))
	}
	return edges
}

func (k *KruskalMST) Weight() float64 {
	return k.weight
}
