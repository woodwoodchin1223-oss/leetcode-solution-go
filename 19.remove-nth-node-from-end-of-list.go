/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    node := head
    prev := head
    total := 0
    for node != nil {
        total += 1
        node = node.Next
    }

    if total == 1 {
        return nil
    }

    count := 1
    node = head
    if total - n + 1 == 1 {
        return head.Next
    }
    for node != nil && count != total - n + 1 {
        prev = node
        node = node.Next
        count += 1
    }
    if node == nil {
        return head
    }
    prev.Next = node.Next
    return head

}
