type Tree struct {
	left *Tree
	right *Tree
	l int
	r int
}

func NewTree(start, end int) *Tree {
	return &Tree{l : start, r : end}
}

func (t *Tree) insert(start, end int) bool {
	curr := t
	for {
		if start >= curr.r {
			if curr.right == nil {
				curr.right = NewTree(start, end)
				return true
			}
			curr = curr.right
		} else if end <= curr.l {
			if curr.left == nil {
				curr.left = NewTree(start, end)
				return true
			}
			curr = curr.left
		} else {
			return false
		}
	}
}

type MyCalendar struct {
    root *Tree
}


func Constructor() MyCalendar {
    return MyCalendar{}
}


func (this *MyCalendar) Book(startTime int, endTime int) bool {
    if this.root == nil {
		this.root = NewTree(startTime, endTime)
		return true
	}
	return this.root.insert(startTime, endTime)
}


/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(startTime,endTime);
 */