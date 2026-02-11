/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    node := &ListNode{}
    head := node
    prev := node
    node1 := l1
    node2 := l2
    tmp := 0
    for ; node1 != nil && node2 != nil; {
        node.Val = ((node1.Val + node2.Val + tmp) % 10)
        tmp = (node1.Val + node2.Val + tmp) / 10
        node1 = node1.Next
        node2 = node2.Next
        node.Next = &ListNode{}
        prev = node
        node = node.Next
    }
    if node1 != nil {
        for ; node1 != nil; {
            node.Val = (node1.Val + tmp) % 10
            tmp = (node1.Val + tmp) / 10
            node1 = node1.Next
            node.Next = &ListNode{}
            prev = node
            node = node.Next
        }
    } else if node2 != nil {
        for ; node2 != nil; {
            node.Val = (node2.Val + tmp) % 10
            tmp = (node2.Val + tmp) / 10
            node2 = node2.Next
            node.Next = &ListNode{}
            prev = node
            node = node.Next
        }
    }
    if tmp != 0 {
        node.Val = tmp
        node.Next = nil
    } else {
        prev.Next = nil
    }
    return head
}
