package tree

var NULL = -1 << 63

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Ints2Tree(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
