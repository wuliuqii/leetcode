package tree

import (
	"strconv"
	"strings"
)

func binaryTreePaths(root *TreeNode) (res []string) {
	var dfs func(node *TreeNode, path []string)
	dfs = func(node *TreeNode, path []string) {
		if node == nil {
			return
		}
		path = append(path, strconv.Itoa(node.Val))
		if node.Left == nil && node.Right == nil {
			res = append(res, strings.Join(path, "->"))
			return
		}
		if node.Left != nil {
			dfs(node.Left, path)
		}
		if node.Right != nil {
			dfs(node.Right, path)
		}
	}

	dfs(root, []string{})
	return
}
