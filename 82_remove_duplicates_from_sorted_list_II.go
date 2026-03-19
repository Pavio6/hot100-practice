package hot100practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// 82. 删除排序链表中的重复元素 II
// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中没有重复出现的数字。
// 示例 1:
// 输入: 1->2->3->3->4->4->5
// 输出: 1->2->5
func DeleteDuplicates02(head *ListNode) *ListNode {
	// 头节点的前一个节点
	dummy := &ListNode{Next: head}
	// 当前节点的前一个节点
	prev := dummy
	// 当前节点
	cur := head

	for cur != nil {
		// 发现重复
		if cur.Next != nil && cur.Val == cur.Next.Val {
			val := cur.Val

			// 跳过所有重复节点
			for cur != nil && cur.Val == val {
				cur = cur.Next
			}

			// 删除整段
			prev.Next = cur
		} else {
			prev = cur
			cur = cur.Next
		}
	}

	return dummy.Next
}
