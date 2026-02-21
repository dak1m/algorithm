package list

type ListNode struct {
	Node int
	Next *ListNode
}

func SimpleBuildListNode(l ...int) *ListNode {
	list := &ListNode{l[0], nil}
	p := list
	for _, v := range l[1:] {
		p.Next = &ListNode{v, nil}
		p = p.Next
	}
	return list
}
