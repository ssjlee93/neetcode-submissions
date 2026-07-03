type ListNode struct {
    Val int
    Next *ListNode
}

type LinkedList struct {
    List []int
    Head *ListNode
    Tail *ListNode
}

func NewLinkedList() *LinkedList {
    return &LinkedList{make([]int, 0), nil, nil}
}

func (ll *LinkedList) Get(index int) int {
    if index < 0 || index > len(ll.List)-1 {
        return -1
    }
    return ll.List[index]
}

func (ll *LinkedList) InsertHead(val int) {
    n := &ListNode{val, ll.Head}
    // set new head
    ll.Head = n
    if ll.Tail == nil {
        ll.Tail = n
    }
    // prepend on slice
    ll.List = append([]int{val}, ll.List...)
}

func (ll *LinkedList) InsertTail(val int) {
    n := &ListNode{val, nil}
    ll.List = append(ll.List, val)
    if ll.Tail == nil {
        ll.Tail = n
        
        if ll.Head == nil {
            ll.Head = n
        }
        return
    }
    ll.Tail.Next = n
    ll.Tail = ll.Tail.Next
}

func (ll *LinkedList) Remove(index int) bool {
    if ll.Head == nil {
        ll.Tail = nil
    }

    if index > len(ll.List)-1 {
        return false
    }

    if index == 0 {
        ll.Head = ll.Head.Next
        ll.List = ll.List[1:]
        return true
    }

    curr := ll.Head
    prev := ll.Head
    for i := 0; i <= index; i++ {
        prev = curr
        curr = curr.Next
    }

    if index == len(ll.List)-1 {
        ll.Tail = prev
        ll.Tail.Next = nil
        ll.List = ll.List[:len(ll.List)-1]
        return true
    }

    prev.Next = curr.Next
    ll.List = append(ll.List[:index], ll.List[index+1:]...)
    return true
}

func (ll *LinkedList) GetValues() []int {
    return ll.List
}
