package main

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) (ans []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		for _, childNode := range node.Children {
			dfs(childNode)
		}
		ans = append(ans, node.Val)
	}
	dfs(root)
	return
}
