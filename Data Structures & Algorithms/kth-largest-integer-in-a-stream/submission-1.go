type KthLargest struct {
    // use heap
    heap []int
    position int
}


func Constructor(k int, nums []int) KthLargest {
    this := KthLargest{heap: []int{0}, position: k}
    // add each element of nums into our heap
    for _, n := range nums {
        this.Add(n)
    }
    return this
}


func (this *KthLargest) Add(val int) int {
    this.heap = append(this.heap, val)
    i := len(this.heap)-1

    for i > 1 && this.heap[i] < this.heap[i/2] {
        this.heap[i], this.heap[i/2] = this.heap[i/2], this.heap[i]
        i = i/2
    }

    if len(this.heap) - 1 > this.position {
        this.pop()
    }

    return this.heap[1]
}

func (this *KthLargest) pop() {
    if len(this.heap) == 1 {
        return
    }
    if len(this.heap) == 2 {
        this.heap = this.heap[:len(this.heap)-1]
        return
    }

    this.heap[1] = this.heap[len(this.heap)-1]
    this.heap = this.heap[:len(this.heap)-1]
    i := 1
    for 2*i < len(this.heap) {
        leftI := 2*i
        rightI := 2*i+1
        // Case : right child exists
        // and right child is smaller than left
        // determine if we should swap with right child
        if rightI < len(this.heap) && this.heap[rightI] < this.heap[leftI] && this.heap[i] > this.heap[rightI] {
            this.heap[i], this.heap[rightI] = this.heap[rightI], this.heap[i]
            i = rightI
        // Case : left child exists
        // determine if we should swap with left child
        } else if this.heap[i] > this.heap[leftI] {
            this.heap[i], this.heap[leftI] = this.heap[leftI], this.heap[i]
            i = leftI
        // Case : no children. We finished our heapify down
        } else {
            break
        }
    }
    return
}