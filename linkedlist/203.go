package linkedlist

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	dummy := &ListNode{}
	dummy.Next = head
	pre := dummy
	cur := dummy
	for cur != nil {
		cur = cur.Next
		for cur != nil && cur.Val == val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = cur
	}

	return dummy.Next
}
