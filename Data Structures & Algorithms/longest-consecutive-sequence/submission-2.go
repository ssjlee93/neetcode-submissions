type UnionFind struct {
    parents map[int]int
    ranks map[int]int
    longest int
}

func New(nums []int) UnionFind {
    uf := UnionFind{
        parents : make(map[int]int),
        ranks : make(map[int]int),
        longest : 1,
    }
    for _, n := range nums {
        uf.parents[n] = n
        uf.ranks[n] = 1
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
    uf.longest = max(uf.longest, uf.ranks[p1], uf.ranks[p2])
    return true
}

func longestConsecutive(nums []int) int {
    if len(nums) <= 1 {
        return len(nums)
    }
    // union find lesson
    uf := New(nums)

    // loop and curr num - 1 joins union
    for _, curr := range nums {
        if _, ok := uf.parents[curr-1]; !ok {
            continue
        }
        uf.union(curr, curr-1)
    }

    return uf.longest
}
