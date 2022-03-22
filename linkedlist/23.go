package linkedlist

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}

	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left + right) / 2
	listLeft := merge(lists, left, mid)
	listRight := merge(lists, mid+1, right)

	return mergeTwoLists(listLeft, listRight)
}
