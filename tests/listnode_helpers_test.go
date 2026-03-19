package tests

import hot100practice "hot100_practice"

func buildLinkedList(nums []int) *hot100practice.ListNode {
	dummy := &hot100practice.ListNode{}
	cur := dummy
	for _, v := range nums {
		cur.Next = &hot100practice.ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func linkedListToSlice(head *hot100practice.ListNode) []int {
	out := []int{}
	for head != nil {
		out = append(out, head.Val)
		head = head.Next
	}
	return out
}

