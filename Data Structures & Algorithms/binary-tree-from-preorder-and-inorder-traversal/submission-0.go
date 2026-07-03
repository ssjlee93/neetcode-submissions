/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    // preorder[0] = always root node
	// find the root in inorder
	// recurse left with preorder 1:mid+1
	// recurse right with preorder mid+1:

	// base case 
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	// define root node
	root := &TreeNode{preorder[0], nil, nil}
	// find mid index
	mid := 0
	for i, n := range inorder {
		if n == preorder[0] {
			mid = i
			break
		}
	}

	// recurse left
	root.Left = buildTree(preorder[1:mid+1], inorder[0:mid])
	// recurse right
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])
	
	return root
}
