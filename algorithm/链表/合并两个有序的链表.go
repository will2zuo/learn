package main

func main() {

}
func mergeTwoLists(h1 *ListNode, h2 *ListNode) *ListNode {
	if h1 == nil {
		return h2
	}
	if h2 == nil {
		return h1
	}
	pre := &ListNode{Value: -1}
	head := pre
	for h1 != nil && h2 != nil {
		if h1.Value < h2.Value {
			pre.Next = h1
			h1 = h1.Next
			pre = pre.Next
		} else {
			pre.Next = h2
			h2 = h2.Next
			pre = pre.Next
		}
	}
	if h1 == nil {
		pre.Next = h2
	}
	if h2 == nil {
		pre.Next = h1
	}
	return head.Next
}
