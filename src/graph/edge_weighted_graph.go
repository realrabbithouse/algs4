package graph

import (
	"algs4/src/typ"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Edge struct {
	v      ID
	w      ID
	weight float64
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) Either() ID {
	return e.v
}

func (e *Edge) Other(vertex ID) (ID, error) {
	if e.v == vertex {
		return e.w, nil
	} else if e.w == vertex {
		return e.v, nil
	} else {
		return 0, errors.New("illegal endpoint")
	}
}

func (e Edge) CompareTo(obj typ.Comparable) int {
	other := obj.(Edge)
	if e.weight < other.weight {
		return -1
	} else if e.weight > other.weight {
		return 1
	} else {
		return 0
	}
}

func (e Edge) String() string {
	return fmt.Sprintf("%d-%d %.2f", e.v, e.w, e.weight)
}

type EdgeWeightedGraph struct {
	V   int
	E   int
	adj [][]Edge
}

func NewEdgeWeightedGraph(V int) (*EdgeWeightedGraph, error) {
	if V < 0 {
		return nil, errors.New("number of vertices must be non-negative")
	}
	return &EdgeWeightedGraph{
		V:   V,
		adj: make([][]Edge, V),
	}, nil
}

func parseSrcDstWeight(line string) (src, dst ID, weight float64, err error) {
	raw := strings.FieldsFunc(line, unicode.IsSpace)
	if len(raw) != 3 {
		err = errors.New("invalid input format in EdgeWeightedGraph constructor")
		return
	}
	parseInt, err := strconv.ParseInt(raw[0], 10, 64)
	if err != nil {
		return
	}
	src = ID(parseInt)
	parseInt, err = strconv.ParseInt(raw[1], 10, 64)
	if err != nil {
		return
	}
	dst = ID(parseInt)
	weight, err = strconv.ParseFloat(raw[2], 64)
	return
}

func ReadEdgeWeightedGraph(scanner *bufio.Scanner) (*EdgeWeightedGraph, error) {
	if scanner == nil {
		return nil, errors.New("argument is nil")
	}
	var graph *EdgeWeightedGraph
	if scanner.Scan() {
		rawV := scanner.Text()
		V, err := strconv.ParseInt(rawV, 10, 64)
		if err != nil {
			return nil, err
		}
		graph, err = NewEdgeWeightedGraph(int(V))
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
			graph.AddEdge(Edge{
				v:      src,
				w:      dst,
				weight: weight,
			})
		} else {
			return nil, errors.New("number of edges is less than the specified")
		}
	}
	return graph, nil
}

func ReadEdgeWeightedGraphFromFile(path string) (*EdgeWeightedGraph, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ReadEdgeWeightedGraph(bufio.NewScanner(f))
}

func (wg *EdgeWeightedGraph) AddEdge(e Edge) {
	err := wg.validateVertex(int(e.v))
	if err != nil {
		return
	}
	err = wg.validateVertex(int(e.w))
	if err != nil {
		return
	}
	wg.adj[e.v] = append(wg.adj[e.v], e)
	wg.adj[e.w] = append(wg.adj[e.w], e)
	wg.E++
}

func (wg *EdgeWeightedGraph) Adj(id int) []Edge {
	err := wg.validateVertex(id)
	if err != nil {
		fmt.Println("Adj:", err)
	}
	return wg.adj[id]
}

func (wg *EdgeWeightedGraph) Degree(id int) int {
	err := wg.validateVertex(id)
	if err != nil {
		fmt.Println("Degree:", err)
		return 0
	}
	return len(wg.adj[id])
}

func (wg *EdgeWeightedGraph) Edges() []Edge {
	edges := make([]Edge, 0, wg.E)
	for i := range wg.adj {
		var selfLoops int
		for j := range wg.adj[i] {
			other, _ := wg.adj[i][j].Other(ID(i))
			if other > ID(i) {
				edges = append(edges, wg.adj[i][j])
			} else if other == ID(i) {
				// add only one copy of each self loop (self loops will be consecutive)
				if selfLoops%2 == 0 {
					edges = append(edges, wg.adj[i][j])
				}
				selfLoops++
			}
		}
	}
	return edges
}

func (wg EdgeWeightedGraph) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(wg.V) + " vertices, " + strconv.Itoa(wg.E) + " edges\n")
	for i := 0; i < wg.V; i++ {
		if len(wg.adj[i]) > 0 {
			builder.WriteString(strconv.Itoa(i) + ": ")
			for j := range wg.adj[i] {
				builder.WriteString(fmt.Sprintf("(%s) ", wg.adj[i][j]))
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (wg *EdgeWeightedGraph) validateVertex(id int) error {
	if id < 0 || id >= wg.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", id, wg.V))
	}
	return nil
}
