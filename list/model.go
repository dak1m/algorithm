package list

type ListNode struct {
	Val  int
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

func SimpleBuildListNodes(list ...[]int) []*ListNode {
	result := make([]*ListNode, 0, len(list))
	for _, l := range list {
		list := &ListNode{l[0], nil}
		p := list
		for _, v := range l[1:] {
			p.Next = &ListNode{v, nil}
			p = p.Next
		}
		result = append(result, list)
	}
	return result
}
