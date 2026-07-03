/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    // recursion
    return divide(lists, 0, len(lists)-1)
}

// divide takes the whole slice and split in half
// imagine the main list has 1,000,000 elements.
// this will divide and conquer the large slice.
func divide(lists []*ListNode, left, right int) *ListNode {
    // if we cross the middle,
    // we exit
    if left > right {
        return nil
    }

    // if they meet right in the middle,
    // we finished all processing
    if left == right {
        return lists[left]
    }

    mid := left + (right-left)/2
    l1 := divide(lists, left, mid)
    l2 := divide(lists, mid+1, right)
    // conquer the subproblems
    // actually merging the ListNodes
    return conquer(l1, l2)
}

func conquer(l1, l2 *ListNode) *ListNode {
    // create n nodes
    dummy := &ListNode{}
    curr := dummy

    // loop until we reach one end
    for l1 != nil && l2 != nil {
        // whichever is smaller attaches to the dummy
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }
    // finish off remainders
    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }
    
    return dummy.Next
}
