package linkedlist

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	reverse := func(l *ListNode) *ListNode {
		var dummy *ListNode
		pre := dummy
		cur := l
		for cur != nil {
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}
		return pre
	}

	rl1 := reverse(l1)
	rl2 := reverse(l2)
	carry := 0
	dummy := &ListNode{}
	cur := dummy
	for rl1 != nil || rl2 != nil || carry != 0 {
		if rl1 != nil {
			carry += rl1.Val
			rl1 = rl1.Next
		}
		if rl2 != nil {
			carry += rl2.Val
			rl2 = rl2.Next
		}
		cur.Next = &ListNode{Val: carry % 10}
		cur = cur.Next
		carry /= 10
	}

	return reverse(dummy.Next)
}
