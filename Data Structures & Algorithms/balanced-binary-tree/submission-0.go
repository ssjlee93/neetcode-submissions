/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    // postorder
	var dfs func(*TreeNode) (int, bool)
	dfs = func(node *TreeNode) (int, bool) {
		if node == nil {
			return 0, true
		}

		l, lok := dfs(node.Left)
		r, rok := dfs(node.Right)
		if !lok || !rok {
			return 999, false
		}
		// check if subtree remains balanced until this point
		balanced := lok && rok && (l-r <= 1 && r-l <= 1)
		// add higher height to total height
		higher := l
		if r > l {
			higher = r
		}
		return 1 + higher, balanced
	}

	_, ok := dfs(root)
	return ok
}
