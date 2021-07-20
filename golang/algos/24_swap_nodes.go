func swapPairs(head *ListNode) *ListNode {
	dummy := new(ListNode)

	dummy.Next = head
	current := dummy

	for current.Next != nil && current.Next.Next != nil {
		first := current.Next
		second := current.Next.Next

		first.Next = second.Next
		current.Next = second
		current.Next.Next = first
		current = current.Next.Next
	}
	return dummy.Next
}
