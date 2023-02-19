/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := ListNode{0, nil}
    cur := &res

    for l1 != nil || l2 != nil {
        sum := 0
        if cur != nil {
            sum = cur.Val
        }
        if l1 != nil {
            sum += l1.Val
            l1 = l1.Next
        }
        if l2 != nil {
            sum += l2.Val
            l2 = l2.Next
        }
        mod := sum % 10
        cur.Val = mod

        if l1 != nil || l2 != nil || sum / 10 > 0 {
            cur.Next = &ListNode{sum / 10, nil}
            cur = cur.Next
        }
    }

    return &res
}