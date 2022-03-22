package linkedlist

func oddEvenList(head *ListNode) *ListNode {
	evenDummy := &ListNode{}
	evenCur := evenDummy
	cur := head
	for cur != nil && cur.Next != nil {
		evenCur.Next = cur.Next
		evenCur = evenCur.Next
		cur.Next = cur.Next.Next
		if cur.Next != nil {
			cur = cur.Next
		}
		evenCur.Next = nil
	}
	if cur != nil {
		cur.Next = evenDummy.Next
	}

	return head
}
