package linkedlist

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	res := &ListNode{Next: head}
	sortPoint, node := head, head.Next
	for node != nil {
		if sortPoint.Val <= node.Val {
			sortPoint = sortPoint.Next
		} else {
			pre := res
			for pre.Next != nil {
				if pre.Next.Val <= node.Val {
					pre = pre.Next
				} else {
					break
				}
			}
			sortPoint.Next = node.Next
			node.Next = pre.Next
			pre.Next = node

		}
		node = sortPoint.Next
	}
	return res.Next
}
