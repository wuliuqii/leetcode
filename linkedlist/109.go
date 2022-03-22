package linkedlist

import . "leetcode/tree"

func sortedListToBST(head *ListNode) *TreeNode {
	var buildTree func(head, tail *ListNode) *TreeNode
	buildTree = func(head, tail *ListNode) *TreeNode {
		if head == tail {
			return nil
		}

		slow, fast := head, head
		for fast != tail && fast.Next != tail {
			slow = slow.Next
			fast = fast.Next.Next
		}

		root := &TreeNode{Val: slow.Val}
		root.Left = buildTree(head, slow)
		root.Right = buildTree(slow.Next, tail)
		return root
	}

	return buildTree(head, nil)
}
