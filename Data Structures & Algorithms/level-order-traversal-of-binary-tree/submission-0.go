/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    // plain BFS
    // at the end of each level, 
    // add to ans
    
    ans := make([][]int, 0)
    
    q := make([]*TreeNode, 0)
    // add root node to q for processing
    if root != nil {
        q = append(q, root)
    }
    
    // begin BFS processing
    for len(q) > 0 {
        // extract lvl size to process
        qSize := len(q)
        lvl := make([]int, 0)

        for i := 0; i < qSize; i++ {
            // pop current node
            node := q[0]
            q = q[1:]

            lvl = append(lvl, node.Val)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }

        // lvl processing finished
        ans = append(ans, lvl)
    }

    return ans
    
}
