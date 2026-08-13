/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

func cloneGraph(node *Node) *Node {
    // catch empty node
    if node == nil {
        return nil
    }

    ans := make(map[*Node]*Node)
    ans[node] = &Node{node.Val, make([]*Node, 0)}
    q := make([]*Node, 0)
    q = append(q, node)

    for len(q) > 0 {
        curr := q[0]
        q = q[1:]

        for _, neighbor := range curr.Neighbors {
            if _, ok := ans[neighbor]; !ok {
                ans[neighbor] = &Node{neighbor.Val, make([]*Node, 0)}
                q = append(q, neighbor)
            }
            ans[curr].Neighbors = append(ans[curr].Neighbors, ans[neighbor])
        }
    }
    return ans[node]
}
