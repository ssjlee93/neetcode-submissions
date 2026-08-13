type Graph struct {
    graph map[int][]int
}

func NewGraph() *Graph {
    graph := make(map[int][]int)
    return &Graph{graph}
}

func (g *Graph) AddEdge(src, dst int) {
    edges, ok := g.graph[src]
    if !ok {
        g.graph[src] = make([]int, 0)
    }
    for _, edge := range edges {
        if edge == dst {
            return
        }
    }
    g.graph[src] = append(g.graph[src], dst)
}

func (g *Graph) RemoveEdge(src, dst int) bool {
    edges, ok := g.graph[src]
    if !ok {
        return false
    }
    for i, edge := range edges {
        if edge == dst {
            g.graph[src] = append(g.graph[src][:i], g.graph[src][i+1:]...)
            return true
        }
    }
    return false
}

func (g *Graph) HasPath(src, dst int) bool {
    visit := make(map[int]bool)
    q := make([]int, 0)
    visit[src] = true
    q = append(q, src)

    for len(q) > 0 {
        curr := q[0]
        q = q[1:]

        if curr == dst {
            return true
        }
        for _, neighbor := range g.graph[curr] {
            if _, ok := visit[neighbor]; !ok {
                visit[neighbor] = true
                q = append(q, neighbor)
            }
        }
    }
    return false
}
