package graph

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Digraph struct {
	V        int
	E        int
	adj      [][]ID
	indegree []int
}

func NewDigraph(V int) (*Digraph, error) {
	if V < 0 {
		return nil, errors.New("number of vertices must be non-negative")
	}
	return &Digraph{
		V:        V,
		adj:      make([][]ID, V),
		indegree: make([]int, V),
	}, nil
}

func ReadDigraph(scanner *bufio.Scanner) (*Digraph, error) {
	if scanner == nil {
		return nil, errors.New("argument is nil")
	}
	var graph *Digraph
	if scanner.Scan() {
		rawV := scanner.Text()
		V, err := strconv.ParseInt(rawV, 10, 64)
		if err != nil {
			return nil, err
		}
		graph, err = NewDigraph(int(V))
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

func ReadDigraphFromFile(path string) (*Digraph, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	return ReadDigraph(bufio.NewScanner(f))
}

func (dg *Digraph) AddEdge(src, dst ID) {
	err := dg.validateVertex(int(src))
	if err != nil {
		return
	}
	err = dg.validateVertex(int(dst))
	if err != nil {
		return
	}
	dg.adj[src] = append(dg.adj[src], dst)
	dg.E++
	dg.indegree[dst]++
}

func (dg *Digraph) Adj(id int) []ID {
	err := dg.validateVertex(id)
	if err != nil {
		fmt.Println("Adj:", err)
		return nil
	}
	return dg.adj[id]
}

func (dg *Digraph) InDegree(id int) int {
	err := dg.validateVertex(id)
	if err != nil {
		fmt.Println("InDegree:", id)
		return 0
	}
	return dg.indegree[id]
}

func (dg *Digraph) OutDegree(id int) int {
	err := dg.validateVertex(id)
	if err != nil {
		fmt.Println("InDegree:", id)
		return 0
	}
	return len(dg.adj[id])
}

func (dg *Digraph) Reverse() *Digraph {
	reverse, _ := NewDigraph(dg.V)
	for i := 0; i < dg.V; i++ {
		for _, dst := range dg.adj[i] {
			reverse.AddEdge(dst, ID(i))
		}
	}
	return reverse
}

func (dg Digraph) String() string {
	var builder strings.Builder
	builder.WriteString(strconv.Itoa(dg.V) + " vertices, " + strconv.Itoa(dg.E) + " edges\n")
	for i := 0; i < dg.V; i++ {
		if len(dg.adj[i]) > 0 {
			builder.WriteString(strconv.Itoa(i) + ": ")
			for _, id := range dg.adj[i] {
				builder.WriteString(fmt.Sprintf("%d ", id))
			}
			builder.WriteString("\n")
		}
	}
	return builder.String()
}

func (dg *Digraph) validateVertex(id int) error {
	if id < 0 || id >= dg.V {
		return errors.New(fmt.Sprintf("vertex %d is not between 0 and %d", id, dg.V))
	}
	return nil
}
