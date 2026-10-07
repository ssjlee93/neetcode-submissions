/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func preorderTraversal(root *TreeNode) []int {
    // iterative DFS lesson
    stack := make([]*TreeNode, 0)
    curr := root
    ans := make([]int, 0)
    for curr != nil || len(stack) > 0 {
        if curr != nil {
            // preorder operation
            ans = append(ans, curr.Val)
            // add right tree to stack
            if curr.Right != nil {
                stack = append(stack, curr.Right)
            }
            // move left 
            curr = curr.Left
        } else {
            curr = stack[len(stack)-1]
            stack = stack[:len(stack)-1]
        }
    }
    return ans
}
