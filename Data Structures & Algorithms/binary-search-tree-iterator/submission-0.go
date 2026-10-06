/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// iterative DFS
// copied solution
type BSTIterator struct {
    flattened []int
    idx int
}

func Constructor(root *TreeNode) BSTIterator {

	bsti := BSTIterator{flattened : make([]int, 0), idx : 0}
    // stack for iterative DFS
    stack := make([]*TreeNode, 0)

    for root != nil || len(stack) > 0 {
        if root != nil {
            stack = append(stack, root)
            root = root.Left
        } else {
            n := len(stack)
            root = stack[n-1]
            stack = stack[:n-1]
            bsti.flattened = append(bsti.flattened, root.Val)
            root = root.Right
        }
    }
    return bsti
}

func (this *BSTIterator) Next() int {
    ans := this.flattened[this.idx]
    this.idx++
    return ans
}

func (this *BSTIterator) HasNext() bool {
	return this.idx < len(this.flattened)
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root)
 * param_1 := obj.Next()
 * param_2 := obj.HasNext()
 */
 