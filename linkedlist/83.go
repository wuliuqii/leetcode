package linkedlist

func deleteDuplicates(head *ListNode) *ListNode {
	pre := head
	cur := head
	for cur != nil {
		for cur != nil && pre.Val == cur.Val {
			cur = cur.Next
		}
		pre.Next = cur
		pre = cur
	}
	return head
}
