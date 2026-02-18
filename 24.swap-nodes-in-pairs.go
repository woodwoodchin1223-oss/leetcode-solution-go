/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    if head.Next == nil {
        return head
    }
    node := head
    next := node.Next
    nnext := node.Next.Next
    nnext = swapPairs(nnext)
    node.Next = nnext
    next.Next = node
    return next
}
