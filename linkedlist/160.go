package linkedlist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tmpA, tmpB := headA, headB
	flagA, flagB := false, false
	for tmpA != nil && tmpB != nil {
		if tmpA == tmpB {
			return tmpA
		}
		tmpA = tmpA.Next
		tmpB = tmpB.Next
		if tmpA == nil && !flagA {
			tmpA = headB
			flagA = true
		}
		if tmpB == nil && !flagB {
			tmpB = headA
			flagB = true
		}
	}
	return nil
}
