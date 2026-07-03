type Heap struct {
    Heap []int
}

func NewHeap() *Heap {
    return &Heap{
        Heap: append([]int{}, 0),
    }
}

func (h *Heap) Push(val int) {
    h.Heap = append(h.Heap, val)
    i := len(h.Heap) - 1

    // Percolate up
    for i > 1 && h.Heap[i] > h.Heap[i/2] {
        // Swap values
        tmp := h.Heap[i]
        h.Heap[i] = h.Heap[i/2]
        h.Heap[i/2] = tmp
        i = i / 2
    }
}

func (h *Heap) Pop() int {
    if len(h.Heap) == 1 {
        return 0
    }
    if len(h.Heap) == 2 {
        popVal := h.Heap[len(h.Heap)-1]
        h.Heap = h.Heap[:len(h.Heap)-1]
        return popVal
    }

    res := h.Heap[1]
    // Move last value to root
    h.Heap[1] = h.Heap[len(h.Heap)-1]
    h.Heap = h.Heap[:len(h.Heap)-1]
    i := 1
    // Percolate down
    for 2*i < len(h.Heap) {
        if 2*i+1 < len(h.Heap) && h.Heap[2*i+1] > h.Heap[2*i] && h.Heap[i] < h.Heap[2*i+1] {
            // Swap right child
            tmp := h.Heap[i]
            h.Heap[i] = h.Heap[2*i+1]
            h.Heap[2*i+1] = tmp
            i = 2*i + 1
        } else if h.Heap[i] < h.Heap[2*i] {
            // Swap left child
            tmp := h.Heap[i]
            h.Heap[i] = h.Heap[2*i]
            h.Heap[2*i] = tmp
            i = 2 * i
        } else {
            break
        }
    }
    return res
}

func (h *Heap) Heapify(arr []int) {
    // 0-th position is moved to the end
    arr = append(arr, arr[0])
    h.Heap = arr
    cur := (len(h.Heap) - 1) / 2

    for cur > 0 {
        // Percolate Down
        i := cur
        for 2*i < len(h.Heap) {
            if 2*i+1 < len(h.Heap) && h.Heap[2*i+1] > h.Heap[2*i] && h.Heap[i] < h.Heap[2*i+1] {
                // Swap right child
                h.Heap[i], h.Heap[2*i+1] = h.Heap[2*i+1], h.Heap[i]
                i = 2*i + 1
            } else if h.Heap[i] < h.Heap[2*i] {
                // Swap left child
                h.Heap[i], h.Heap[2*i] = h.Heap[2*i], h.Heap[i]
                i = 2 * i
            } else {
                break
            }
        }
        cur--
    }
}

func lastStoneWeight(stones []int) int {
    h := NewHeap()
    h.Heapify(stones)

    for len(h.Heap) > 2 {
        stone1 := h.Pop()
        stone2 := h.Pop()

        diff := stone1 - stone2
        if diff < 0 {
            diff *= -1
        }

        if diff > 0 {
            h.Push(diff)
        }
    }

    if len(h.Heap) == 1 {
        return 0
    }
    return h.Pop()
}
