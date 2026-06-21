/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    head = skip(head)
    result, left := head, head
    for head != nil && head.Next != nil {
        if head.Next.Val != head.Val {
            left = head
        } else {
            left.Next = skip(head)
        }
        head = head.Next
    }
    return result
}

func skip(head *ListNode) *ListNode {
    for head.Next != nil && head.Val == head.Next.Val {
        val := head.Val
        for head.Val == val {
            head = head.Next
            if head == nil {
                return head
            }
        }
    }
    return head
}
