/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    i := 0
    ans := 0
    var dfs func(*TreeNode)
    dfs = func(node *TreeNode)  {
        if node == nil {
            return 
        }

        dfs(node.Left)
        i++
        if i == k {
            ans = node.Val
        }
        dfs(node.Right)
        return
    }
    dfs(root)
    return ans
        
}
