package hot100practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 141. 环形链表
func HasCycle(head *ListNode) bool {
    // 使用快慢指针
    if head == nil || head.Next == nil {
        return false
    }

    slow := head
    fast := head

    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next

        if slow == fast {
            return true
        }
    }

    return false
}