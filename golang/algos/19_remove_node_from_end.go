/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if n == 0 || head == nil {
		return nil
	}
	res := new(ListNode)
	res.Next = head
	head = res
	p, q := head, head
	for i := 0; i <= n; i++ {
		if q != nil {
			q = q.Next
		} else {
			return nil
		}
	}
	for q != nil {
		p = p.Next
		q = q.Next
	}
	p.Next = p.Next.Next
	return head.Next
}
