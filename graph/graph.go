package graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type ID int

type Graph struct {
	V   int
	E   int
	adj [][]ID
}

func NewGraph(V int) (*Graph, error) {
	if V < 0 {
		return nil, errors.New("number of vertices must be non-negative")
	}
	return &Graph{
		V:   V,
		adj: make([][]ID, V),
	}, nil
}

func parseSrcDst(line string) (src, dst ID, err error) {
	raw := strings.FieldsFunc(line, unicode.IsSpace)
	if len(raw) != 2 {
		err = errors.New("invalid input format in Graph constructor")
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
	return
}

func ReadGraph(scanner *bufio.Scanner) (*Graph, error) {
	if scanner == nil {
		return nil, errors.New("argument is nil")
	}
	var graph *Graph
	if scanner.Scan() {
		rawV := scanner.Text()
		V, err := strconv.ParseInt(rawV, 10, 64)
		if err != nil {
			return nil, err
		}
		graph, err = NewGraph(int(V))
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
			src, dst, err := parseSrcDst(line)
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
			graph.AddEdge(src, dst)
		} else {
			return nil, errors.New("number of edges is less than the specified")
		}
	}
	return graph, nil
}

func ReadGraphFromFile(path string) (*Graph, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ReadGraph(bufio.NewScanner(f))
}

func (g *Graph) AddEdge(src, dst ID) {
	err := g.validateVertex(int(src))
	if err != nil {
		return
	}
	err = g.validateVertex(int(dst))
	if err != nil {
		return
	}
	g.adj[src] = append(g.adj[src], dst)
	g.adj[dst] = append(g.adj[dst], src)
	g.E++
}

func (g *Graph) Adj(id int) []ID {
	err := g.validateVertex(id)
	if err != nil {
		fmt.Println("Adj:", err)
		return nil
	}
	return g.adj[id]
}

func (g *Graph) Degree(id int) int {
	err := g.validateVertex(id)
	if err != nil {
		fmt.Println("Degree:", err)
		return 0
	}
	return len(g.adj[id])
}

func (g Graph) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(g.V) + " vertices, " + strconv.Itoa(g.E) + " edges\n")
	for i := 0; i < g.V; i++ {
		if len(g.adj[i]) > 0 {
			builder.WriteString(strconv.Itoa(i) + ": ")
			for _, id := range g.adj[i] {
				builder.WriteString(fmt.Sprintf("%d ", id))
			}
		}
		builder.WriteString("\n")
	}
	return builder.String()
}

func (g *Graph) validateVertex(id int) error {
	if id < 0 || id >= g.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", id, g.V))
	}
	return nil
}
