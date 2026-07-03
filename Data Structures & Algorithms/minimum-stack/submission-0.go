type MinStack struct {
	stack []int
	smallz []int
}

func Constructor() MinStack {
	return MinStack{make([]int, 0), make([]int, 0)}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)

	// when smallest stack is empty, shove it in
	if len(this.smallz) == 0 {
		this.smallz = append(this.smallz, 0)
		return
	}
	
	// top of smallest stack == indext of smallest elment in main stack
	smallestIdx := this.smallz[len(this.smallz)-1]
	smallest := this.stack[smallestIdx]
	// then we push smaller index between the top and curr val into smallz
	if val < smallest {
		smallestIdx = len(this.stack)-1
	}
	this.smallz = append(this.smallz, smallestIdx)
	return
}

func (this *MinStack) Pop() {
	// pop from main stack
	this.stack = this.stack[:len(this.stack)-1]
	// we also pop from the smallest indices stack
	this.smallz = this.smallz[:len(this.smallz)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.stack[this.smallz[len(this.smallz)-1]]
}
