package hot100practice

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // 876. 链表的中间节点
func MiddleNode(head *ListNode) *ListNode {
    slow := head
    fast := head
    // [1, 2, 3, 4, 5]
    // 2 - 3
    // 3 - 5
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }
    return slow
}