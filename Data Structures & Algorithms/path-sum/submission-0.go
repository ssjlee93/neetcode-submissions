/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func hasPathSum(root *TreeNode, targetSum int) bool {

    var dfs func(*TreeNode, int) bool
    dfs = func(node *TreeNode, sum int) bool {
        if node == nil {
            return false
        }

        sum += node.Val

        // root-to-leaf : if leaf, check for our answer
        if node.Left == nil && node.Right == nil {
            ans := sum == targetSum
            return ans
        }
        left := dfs(node.Left, sum)
        right := dfs(node.Right, sum)
        
        return left || right
    }
    sum := 0
    return dfs(root, sum)

}

