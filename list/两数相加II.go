package list

/*
https://leetcode.cn/problems/add-two-numbers-ii/description/

给你两个 非空 链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储一位数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例1：
输入：l1 = [7,2,4,3], l2 = [5,6,4]
输出：[7,8,0,7]

示例2：
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[8,0,7]

示例3：
输入：l1 = [0], l2 = [0]
输出：[0]


提示：

链表的长度范围为 [1, 100]
0 <= node.val <= 9
输入数据保证链表代表的数字无前导 0


进阶：如果输入链表不能翻转该如何解决？
*/

func addTwoNumbersII(l1 *ListNode, l2 *ListNode) *ListNode {
	stack1, stack2 := make([]int, 0), make([]int, 0)
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}
	carry := 0
	dummy := &ListNode{-1, nil}
	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		val := carry
		if len(stack1) > 0 {
			p1 := stack1[len(stack1)-1]

			stack1 = stack1[:len(stack1)-1]
			val += p1
		}
		if len(stack2) > 0 {
			p2 := stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
			val += p2
		}
		carry = val / 10
		val = val % 10
		node := &ListNode{val, dummy.Next}
		dummy.Next = node
	}

	return dummy.Next
}
