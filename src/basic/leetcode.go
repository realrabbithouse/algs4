package basic

// 785. Is Graph Bipartite?

type Color bool

type Bipartite struct {
	graph       [][]int
	visited     []bool
	colors      []Color
	isBipartite bool
}

func (b *Bipartite) dfs(u int) {
	b.visited[u] = true
	for _, v := range b.graph[u] {
		if !b.visited[v] {
			b.colors[v] = !b.colors[u]
			b.dfs(v)
		} else if b.colors[v] == b.colors[u] {
			b.isBipartite = false
			break
		}
	}
}

func isBipartite(graph [][]int) bool {
	n := len(graph)
	b := Bipartite{
		graph:       graph,
		visited:     make([]bool, n),
		colors:      make([]Color, n),
		isBipartite: true,
	}
	for i := 0; i < n; i++ {
		if !b.visited[i] && b.isBipartite {
			b.dfs(i)
		}
	}
	return b.isBipartite
}
