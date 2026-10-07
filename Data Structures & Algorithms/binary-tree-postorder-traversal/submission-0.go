/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func postorderTraversal(root *TreeNode) []int {
    // iterative DFS lesson
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	visit := make([]bool, 0)
	visit = append(visit, false)
	ans := make([]int, 0)

	for len(stack) > 0 {
		// pull out curr node to process and its visitation
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visited := visit[len(visit)-1]
		visit = visit[:len(visit)-1]

		// make sure curr isn't nil
		if curr != nil {
			// visited already : process the node
			if visited {
				ans = append(ans, curr.Val)
			} else {
				// unvisited : insert curr node as visited and add its children as unvisited
				stack = append(stack, curr, curr.Right, curr.Left)
				visit = append(visit, true, false, false)
			}
		}
	}
	return ans
}
