package linkedlist

func reorderList(head *ListNode) {
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

	if head == nil || head.Next == nil {
		return
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	left := head
	right := reverse(slow.Next)
	slow.Next = nil

	for left != nil && right != nil {
		leftNext := left.Next
		rightNext := right.Next
		left.Next = right
		right.Next = leftNext
		left = leftNext
		right = rightNext
	}
}
