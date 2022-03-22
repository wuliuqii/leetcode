package linkedlist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	reverse := func(node *ListNode) *ListNode {
		var prev *ListNode
		curr := node

		for curr != nil {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}

		return prev
	}

	if left == right {
		return head
	}

	dummy := &ListNode{}
	dummy.Next = head

	leftPrev := dummy
	for i := 1; i < left; i++ {
		leftPrev = leftPrev.Next
	}
	rightNode := dummy
	for i := 1; i <= right; i++ {
		rightNode = rightNode.Next
	}

	leftNode := leftPrev.Next
	rightNext := rightNode.Next
	rightNode.Next = nil
	leftPrev.Next = reverse(leftNode)
	leftNode.Next = rightNext

	return dummy.Next
}
