package linkedlist

func deleteDuplicates2(head *ListNode) *ListNode {
	pHead := &ListNode{Next: head}
	head = pHead
	for head.Next != nil && head.Next.Next != nil {
		if head.Next.Val == head.Next.Next.Val {
			tmp := head.Next.Val
			for head.Next != nil && head.Next.Val == tmp {
				head.Next = head.Next.Next
			}
		} else {
			head = head.Next
		}
	}
	return pHead.Next
}
