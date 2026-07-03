/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    // recursion version
    if head == nil {
        return nil
    }

    prev := head
    if head.Next != nil {
        prev = reverseList(head.Next)
        head.Next.Next = head
    }
    head.Next = nil 

    return prev
}
