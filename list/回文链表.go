package list

import "reflect"

/*
给你一个单链表的头节点 head ，请你判断该链表是否为回文链表。如果是，返回 true ；否则，返回 false 。

示例 1：


输入：head = [1,2,2,1]
输出：true
示例 2：


输入：head = [1,2]
输出：false
提示：

链表中节点数目在范围[1, 105] 内
0 <= Node.val <= 9
进阶：你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

题目来源：力扣 234. 回文链表。https://leetcode.cn/problems/palindrome-linked-list/
*/

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	newList := reverseList(head)

	if reflect.DeepEqual(head, newList) {
		return true
	}

	return false
}

func reverseList(head *ListNode) *ListNode {
	// 反转链表，与原链表做判断
	var newList *ListNode
	cur := head
	for cur != nil {
		newNode := &ListNode{
			Val:  cur.Val,
			Next: newList,
		}
		newList = newNode
		cur = cur.Next
	}

	return newList
}
