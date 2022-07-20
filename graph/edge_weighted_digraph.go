package graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DirectedEdge struct {
	from   ID
	to     ID
	weight float64
}

func (e *DirectedEdge) Src() ID {
	return e.from
}

func (e *DirectedEdge) Dst() ID {
	return e.to
}

func (e *DirectedEdge) Weight() float64 {
	return e.weight
}

func (e DirectedEdge) String() string {
	return fmt.Sprintf("%d->%d %.8f", e.from, e.to, e.weight)
}

type EdgeWeightedDigraph struct {
	V        int
	E        int
	adj      [][]DirectedEdge
	indegree []int
}

func NewEdgeWeightedDigraph(V int) (*EdgeWeightedDigraph, error) {
	if V < 0 {
		return nil, errors.New("number of vertices must be non-negative")
	}
	return &EdgeWeightedDigraph{
		V:        V,
		adj:      make([][]DirectedEdge, V),
		indegree: make([]int, V),
	}, nil
}

func ReadEdgeWeightedDigraph(scanner *bufio.Scanner) (*EdgeWeightedDigraph, error) {
	if scanner == nil {
		return nil, errors.New("argument is nil")
	}
	var graph *EdgeWeightedDigraph
	if scanner.Scan() {
		rawV := scanner.Text()
		V, err := strconv.ParseInt(rawV, 10, 64)
		if err != nil {
			return nil, err
		}
		graph, err = NewEdgeWeightedDigraph(int(V))
		if err != nil {
			return nil, err
		}
	}
	var E int
	if scanner.Scan() {
		rawE := scanner.Text()
		parseE, err := strconv.ParseInt(rawE, 10, 64)
		if err != nil {
			return nil, err
		}
		E = int(parseE)
		if E < 0 {
			return nil, errors.New("number of edges must be non-negative")
		}
	}
	for i := 0; i < E; i++ {
		if scanner.Scan() {
			line := scanner.Text()
			src, dst, weight, err := parseSrcDstWeight(line)
			if err != nil {
				return nil, err
			}
			err = graph.validateVertex(int(src))
			if err != nil {
				return nil, err
			}
			err = graph.validateVertex(int(dst))
			if err != nil {
				return nil, err
			}
			graph.AddEdge(DirectedEdge{
				from:   src,
				to:     dst,
				weight: weight,
			})
		} else {
			return nil, errors.New("number of edges is less than the specified")
		}
	}
	return graph, nil
}

func ReadEdgeWeightedDigraphFromFile(path string) (*EdgeWeightedDigraph, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ReadEdgeWeightedDigraph(bufio.NewScanner(f))
}

func (edg *EdgeWeightedDigraph) AddEdge(e DirectedEdge) {
	err := edg.validateVertex(int(e.from))
	if err != nil {
		return
	}
	err = edg.validateVertex(int(e.to))
	if err != nil {
		return
	}
	edg.adj[e.from] = append(edg.adj[e.from], e)
	edg.E++
}

func (edg *EdgeWeightedDigraph) Adj(id int) []DirectedEdge {
	err := edg.validateVertex(id)
	if err != nil {
		fmt.Println("Adj:", err)
		return nil
	}
	return edg.adj[id]
}

func (edg *EdgeWeightedDigraph) InDegree(id int) int {
	err := edg.validateVertex(id)
	if err != nil {
		fmt.Println("InDegree:", id)
		return 0
	}
	return edg.indegree[id]
}

func (edg *EdgeWeightedDigraph) OutDegree(id int) int {
	err := edg.validateVertex(id)
	if err != nil {
		fmt.Println("InDegree:", id)
		return 0
	}
	return len(edg.adj[id])
}

func (edg *EdgeWeightedDigraph) Reverse() *EdgeWeightedDigraph {
	reverse, _ := NewEdgeWeightedDigraph(edg.V)
	for i := 0; i < edg.V; i++ {
		for j := range edg.adj[i] {
			e := edg.adj[i][j]
			reverse.AddEdge(DirectedEdge{
				from:   e.to,
				to:     e.from,
				weight: e.weight,
			})
		}
	}
	return reverse
}

func (edg EdgeWeightedDigraph) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(edg.V) + " vertices, " + strconv.Itoa(edg.E) + " edges\n")
	for i := 0; i < edg.V; i++ {
		if len(edg.adj[i]) > 0 {
			builder.WriteString(strconv.Itoa(i) + ": ")
			for j := range edg.adj[i] {
				builder.WriteString(fmt.Sprintf("(%s) ", edg.adj[i][j]))
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (edg *EdgeWeightedDigraph) validateVertex(id int) error {
	if id < 0 || id >= edg.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", id, edg.V))
	}
	return nil
}
