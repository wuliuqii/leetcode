package tree

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	var level []*Node
	level = append(level, root)
	for len(level) != 0 {
		var tmp []*Node
		for i, node := range level {
			if i == len(level)-1 {
				node.Next = nil
			} else {
				node.Next = level[i+1]
			}
			if node.Left != nil {
				tmp = append(tmp, node.Left)
			}
			if node.Right != nil {
				tmp = append(tmp, node.Right)
			}
		}
		level = tmp
	}
	return root
}
