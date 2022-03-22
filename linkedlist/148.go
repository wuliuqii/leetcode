package linkedlist

func sortList(head *ListNode) *ListNode {
	return sortPart(head, nil)
}

func sortPart(head *ListNode, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}

	return mergeTwoLists(sortPart(head, slow), sortPart(slow, tail))
}
