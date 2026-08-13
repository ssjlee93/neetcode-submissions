/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// this is undirected and connected
// so one entry node leads to the entire node.
func cloneGraph(node *Node) *Node {
    ans := make(map[*Node]*Node)

    var dfs func(*Node) *Node
    dfs = func(node *Node) *Node {
        // base case
        if node == nil {
            return nil
        }
        // visited
        if _, ok := ans[node]; ok {
            return ans[node]
        }

        clone := &Node{Val: node.Val}

        ans[node] = clone
        for _, neighbor := range node.Neighbors {
            clone.Neighbors = append(clone.Neighbors, dfs(neighbor))
        }
        return clone
    }
    return dfs(node)
}
