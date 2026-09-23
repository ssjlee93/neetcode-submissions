/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
    // fast slow pointer lesson
	// flip fast pointer until it hits end
	// initialize fast pointer on next so it lands on n/2 -1
	
	// edge : min len lower than 2 -> handled by constraints
	fast, slow := head, head
	var prev *ListNode
	prev = nil
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// found mdi at slow

	// flip rest of nodes backwards
	for slow != nil {
		tmp := slow.Next
		slow.Next = prev
		prev = slow
		slow = tmp
	}
	
	// move from head and prev to match the twins
	ans := 0
	pointer := head
	for prev != nil {
		curr := prev.Val + pointer.Val
		ans = max(ans, curr)
		prev = prev.Next
		pointer = pointer.Next
	}
	return ans
}