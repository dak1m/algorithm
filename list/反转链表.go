package list

/*
https://leetcode.cn/problems/reverse-linked-list/
给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。

示例 1：
输入：head = [1,2,3,4,5]
输出：[5,4,3,2,1]

示例 2：
输入：head = [1,2]
输出：[2,1]

示例 3：
输入：head = []
输出：[]
*/
func reverseListIteration(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var pre, cur, next *ListNode
	pre, cur, next = nil, head, head.Next
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = next
		if next != nil {
			next = next.Next
		}
	}

	return pre
}

func reverseListRecursion(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	last := reverseListRecursion(head.Next)
	head.Next.Next = head

	head.Next = nil

	return last
}
