
// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
// flipped the comparator to make it a max heap.
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Remove(i int) any { return i }

func lastStoneWeight(stones []int) int {
	s := IntHeap(stones)
    h := &s
	heap.Init(h)

	fmt.Printf("maximum: %d\n", (*h)[0])
	for h.Len() > 1 {
        stone1 := heap.Pop(h).(int)
        stone2 := heap.Pop(h).(int)
        diff := stone1 - stone2
        if diff < 0 {
            diff *= -1
        }
        if diff > 0 {
            heap.Push(h, diff)
        }
	}

        // Check if heap is empty
    if h.Len() == 0 {
        return 0
    }
    return heap.Pop(h).(int)
}
