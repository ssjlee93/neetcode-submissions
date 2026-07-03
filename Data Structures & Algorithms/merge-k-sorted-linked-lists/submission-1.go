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

    return conquer(l1, l2)
}

func conquer(l1, l2 *ListNode) *ListNode {
    dummy := &ListNode{}
    curr := dummy

    for l1 != nil && l2 != nil {
        if l1.Val <= l2.Val {
            curr.Next = l1
            l1 = l1.Next
        } else {
            curr.Next = l2
            l2 = l2.Next
        }
        curr = curr.Next
    }

    if l1 != nil {
        curr.Next = l1
    } else {
        curr.Next = l2
    }

    return dummy.Next
}
