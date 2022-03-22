package linkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func Ints2List(nums []int) *ListNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	head := &ListNode{Val: nums[0]}
	cur := head
	for i := 1; i < n; i++ {
		cur.Next = &ListNode{Val: nums[i]}
		cur = cur.Next
	}
	return head
}

// Ints2ListWithCycle returns a list whose tail point to pos-indexed node
// head's index is 0
// if pos = -1, no cycle
func Ints2ListWithCycle(nums []int, pos int) *ListNode {
	head := Ints2List(nums)
	if pos == -1 {
		return head
	}
	c := head
	for pos > 0 {
		c = c.Next
		pos--
	}
	tail := c
	for tail.Next != nil {
		tail = tail.Next
	}
	tail.Next = c
	return head
}
