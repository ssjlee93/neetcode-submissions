type UnionFind struct {
    parents map[int]int
    ranks map[int]int
}

func Constructor(n int) UnionFind {
    uf := UnionFind{ parents : make(map[int]int), ranks : make(map[int]int)}
    for i := range n+1 {
        uf.parents[i] = i
        uf.ranks[i] = 0
    }
    return uf
}

func (uf *UnionFind) find(n int) int {
    p := uf.parents[n]
    for p != uf.parents[p] {
        uf.parents[p] = uf.parents[uf.parents[p]]
        p = uf.parents[p]
    }
    return p
}

func (uf *UnionFind) union(n1, n2 int) bool {
    p1, p2 := uf.find(n1), uf.find(n2)
    if p1 == p2 {
        return false
    }

    if uf.ranks[p1] > uf.ranks[p2] {
        uf.parents[p2] = p1
    } else if uf.ranks[p2] > uf.ranks[p1] {
        uf.parents[p1] = p2
    } else {
        uf.parents[p1] = p2
        uf.ranks[p2]++
    }
    return true
}

func findRedundantConnection(edges [][]int) []int {
    // union find lesson

    // edges need to be formed as UnionFind
    uf := Constructor(len(edges))

    for _, edge := range edges {
        if !uf.union(edge[0], edge[1]) {
            return edge
        }
    }
    return []int{}
}
