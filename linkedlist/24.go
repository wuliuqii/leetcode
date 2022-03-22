package linkedlist

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}

	pre := dummy
	for head != nil && head.Next != nil {
		tmp := head.Next.Next
		pre.Next = head.Next
		head.Next = nil
		pre.Next.Next = head
		head = tmp
		pre = pre.Next.Next
	}
	if head != nil {
		pre.Next = head
	}

	return dummy.Next
}
