package hot100practice

type ListNode struct {
	Val  int
	Next *ListNode
}

// 83. 删除排序链表中的重复元素
func DeleteDuplicates(head *ListNode) *ListNode {

	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}
