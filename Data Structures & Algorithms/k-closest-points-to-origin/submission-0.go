// An IntHeap is a min-heap of ints.
type IntHeap [][3]int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i][0] < h[j][0] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([3]int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
    // use min heap to track euclidian
    h := &IntHeap{}
    heap.Init(h)
    // for each point, calculate euclidian
    for i := 0; i < len(points); i++ {
        x, y := points[i][0], points[i][1]
        d := calcEuclidean(x, y)
        // push euclidian, x, y as an element
        el := [3]int{d, x, y}
        // push to min heap
        heap.Push(h, el)
    }
    ans := [][]int{}
    // when pulling k num of values
    for i := 0; i < k; i++ {
        el := heap.Pop(h).([3]int)
        ans = append(ans, []int{el[1], el[2]})
    }
    // pop k times
    return ans
}

func calcEuclidean(x, y int) int {
    return x*x + y*y
}
