/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
    result := &ListNode{}
    right := &ListNode{}
    first := right
    left := result
    for head != nil {
        next := head.Next
        if head.Val < x {
            head.Next = nil
            left.Next = head
            left = left.Next
        } else {
            head.Next = nil
            right.Next = head
            right = right.Next
        }
        head = next
    }
    left.Next = first.Next
    return result.Next
}