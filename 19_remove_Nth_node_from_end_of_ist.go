package hot100practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 
// 19. 删除链表的倒数第 N 个节点
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
    // dummy相当于在原链表前加一个前驱节点
    dummy := &ListNode{Next: head}
    
    slow := dummy
    fast := dummy

    // fast 先走 n 步
    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    // 一起走
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }

    // 删除节点
    slow.Next = slow.Next.Next

    return dummy.Next

}