package linked_list

// ListNode 链表遍历
type ListNode struct {
	val  int
	next *ListNode
}

func traverse(head *ListNode) {
	for p := head; p != nil; p = p.next {
		// 迭代访问 p.val
	}
}

func recursiveTraverse(head *ListNode) {
	// 递归访问 head.val
	recursiveTraverse(head.next)
}
