type UnionFind struct {
	parents map[int]int
	ranks map[int]int
	numComponents int
}

func NewUnionFind(n int) *UnionFind {
	uf := UnionFind{parents : make(map[int]int), ranks : make(map[int]int), numComponents : n}
	for i := 0; i <= n; i++ {
		uf.parents[i] = i
		uf.ranks[i] = 0
	}
	return &uf
}

func (uf *UnionFind) Find(x int) int {
	p := uf.parents[x]
	for p != uf.parents[p] {
		uf.parents[p] = uf.parents[uf.parents[p]]
		p = uf.parents[p]
	}
	return p
}

func (uf *UnionFind) IsSameComponent(x, y int) bool {
	p1, p2 := uf.Find(x), uf.Find(y)
	return p1 == p2
}

func (uf *UnionFind) Union(x, y int) bool {
	p1, p2 := uf.Find(x), uf.Find(y)
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
	uf.numComponents--
	return true
}

func (uf *UnionFind) GetNumComponents() int {
	return uf.numComponents
}
