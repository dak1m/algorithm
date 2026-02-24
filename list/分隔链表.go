package list

/*
86. 分隔链表 https://leetcode.cn/problems/partition-list/
给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。

你应当 保留 两个分区中每个节点的初始相对位置。

示例 1：
输入：head = [1,4,3,2,5,2], x = 3
输出：[1,2,2,4,3,5]

示例 2：
输入：head = [2,1], x = 2
输出：[1,2]

提示：
链表中节点的数目在范围 [0, 200] 内
-100 <= Node.val <= 100
-200 <= x <= 200
*/

func partition(head *ListNode, x int) *ListNode {
	d1 := &ListNode{-1, nil} // 虚拟头节点
	d2 := &ListNode{-1, nil} // 虚拟头节点

	p1, p2 := d1, d2 // 作用：移动节点需要copy一个变量
	p := head        // 同理
	for p != nil {
		if p.Val < x {
			p1.Next = p
			p1 = p1.Next
		} else {
			p2.Next = p
			p2 = p2.Next
		}
		tmp := p.Next
		p.Next = nil // 断开p后面的指针
		p = tmp
	}

	p1.Next = d2.Next

	return d1.Next
}
