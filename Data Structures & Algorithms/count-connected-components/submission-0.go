type UnionFind struct {
    parents map[int]int
    ranks map[int]int
    numComponents int
}

func New(n int) UnionFind {
    uf := UnionFind{
        parents : make(map[int]int),
        ranks : make(map[int]int),
        numComponents : 0,
    }
    for i := 0; i < n; i++ {
        uf.parents[i] = i
        uf.ranks[i] = 0
        uf.numComponents++
    }
    return uf
}

func (uf *UnionFind) find(x int) int {
    p := uf.parents[x]
    for p != uf.parents[p] {
        uf.parents[p] = uf.parents[uf.parents[p]]
        p = uf.parents[p]
    }
    return p
}

func (uf *UnionFind) union(x, y int) bool {
    p1, p2 := uf.find(x), uf.find(y)
    if p1 == p2 {
        return false
    }

    if uf.ranks[p1] > uf.ranks[p2] {
        uf.parents[p2] = p1
        uf.ranks[p1] += uf.ranks[p2]
    } else if uf.ranks[p2] > uf.ranks[p1] {
        uf.parents[p1] = p2
        uf.ranks[p2] += uf.ranks[p1]
    } else {
        uf.parents[p1] = p2
        uf.ranks[p2] += uf.ranks[p1]
    }
    uf.numComponents--
    return true
}

func countComponents(n int, edges [][]int) int {
    // Union Find lesson
    // this is similar to the counting num components problem.
    // build union find on 0 ~ n-1. 
    uf := New(n)
    // union each edge
    // each beginning of edge is a parent.
    for _, edge := range edges {
        uf.union(edge[0], edge[1])
    }

    // get numcomponents
    return uf.numComponents
}
