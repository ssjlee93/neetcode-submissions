/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    ans := make([]int, 0)
    return inorder(root, ans)
}

func inorder(root *TreeNode, ans []int) []int {
    if root == nil {
        return ans
    }
    ans = inorder(root.Left, ans)
    ans = append(ans, root.Val)
    ans = inorder(root.Right, ans)

    return ans
} 
