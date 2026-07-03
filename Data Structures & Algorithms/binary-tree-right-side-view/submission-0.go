/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    // plain BFS
	// but only add right or left node
	// add right first
	// add left if right doesn't exist and left exists

	ans := make([]int, 0)
	
	// create queue for BFS
	q := make([]*TreeNode, 0)
	// add root
	if root != nil {
		q = append(q, root)
	}
	// process BFS
	for len(q) > 0 {
		// extract just the lvl elements
		lvl := len(q)

		ans = append(ans, q[0].Val)
		
		for i := 0; i < lvl; i++ {
			// pop node
			node := q[0]
			q = q[1:]

			if node.Right != nil {
				q = append(q, node.Right)
			} 
			if node.Left != nil {
				q = append(q, node.Left)
			}
			
		}
	}
	return ans
}
