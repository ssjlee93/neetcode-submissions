type MyStack struct {
    S []int
}

func Constructor() MyStack {
    return MyStack{make([]int, 0)}
}

func (this *MyStack) Push(x int) {
    this.S = append(this.S, x)
    return
}

func (this *MyStack) Pop() int {
    popped := this.S[len(this.S)-1]
    this.S = this.S[:len(this.S)-1]
    return popped
}

func (this *MyStack) Top() int {
    return this.S[len(this.S)-1]
}

func (this *MyStack) Empty() bool {
    return len(this.S) == 0
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param2 := obj.Pop();
 * param3 := obj.Top();
 * param4 := obj.Empty();
 */
