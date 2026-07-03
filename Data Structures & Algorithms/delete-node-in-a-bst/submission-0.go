/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
	// base case
	// when we reach a leaf, 
	// we return nil so we can use this obj to delete
	if root == nil {
		return nil 
	}

	// traverse either side until we find our target
	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else { // target found !

		// if either node is missing, 
		// return the other side
		// the other side will be assigned in place of target node
		// by the code above
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			// if we have 2 children, 
			// we find the min node. 
			// min node will always be a leaf node.
			minVal := findMin(root.Right)
			// replace the value of min w our target
			root.Val = minVal.Val
			// find the original of min node, 
			// return nil when we find it. 
			// then assign nil in its place
			// by the code above
			root.Right = deleteNode(root.Right, minVal.Val)
		}

	}
	// if all operations succeeded, 
	// return the current node to finish the call stack
    return root
}

func findMin(root *TreeNode) *TreeNode {
	curr := root
	for curr != nil && curr.Left != nil {
		curr = curr.Left
	}
	return curr
}