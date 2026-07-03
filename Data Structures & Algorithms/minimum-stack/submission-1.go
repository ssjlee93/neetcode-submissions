type MinStack struct {
	stack [][2]int
}

func Constructor() MinStack {
	return MinStack{make([][2]int, 0)}
}

func (this *MinStack) Push(val int) {
	pair := [2]int{val, val}
	// when smallest stack is empty, shove it in
	if len(this.stack) == 0 {
		this.stack = append(this.stack, pair)
		return
	}
	
	// top of smallest stack == indext of smallest elment in main stack
	lastItem := this.stack[len(this.stack)-1]
	// then we push smaller index between the top and curr val into smallz
	if val < lastItem[1] {
		pair[1] = val
	} else {
		pair[1] = lastItem[1]
	}
	this.stack = append(this.stack, pair)
	return
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1][0]
}

func (this *MinStack) GetMin() int {
	return this.stack[len(this.stack)-1][1]
}
