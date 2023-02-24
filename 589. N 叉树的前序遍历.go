package main

// Node
/** Definition for a Node.
* type Node struct {
*     Val int
*     Children []*Node
* }
 */
type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val)
		for _, childrenNode := range node.Children {
			dfs(childrenNode)
		}
	}
	dfs(root)
	return
}
