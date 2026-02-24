package list

import "container/heap"

/*
23. 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/
给你一个链表数组，每个链表都已经按升序排列。
请你将所有链表合并到一个升序链表中，返回合并后的链表。

示例 1：
输入：lists = [[1,4,5],[1,3,4],[2,6]]
输出：[1,1,2,3,4,4,5,6]
解释：链表数组如下：
[
  1->4->5,
  1->3->4,
  2->6
]
将它们合并到一个有序链表中得到。
1->1->2->3->4->4->5->6

示例 2：
输入：lists = []
输出：[]
示例 3：

输入：lists = [[]]
输出：[]

提示：
k == lists.length
0 <= k <= 10^4
0 <= lists[i].length <= 500
-10^4 <= lists[i][j] <= 10^4
lists[i] 按 升序 排列
lists[i].length 的总和不超过 10^4
*/

type PriorityQueue []*ListNode

var _ heap.Interface = (*PriorityQueue)(nil)

func (p *PriorityQueue) Len() int {
	return len(*p)
}

func (p *PriorityQueue) Less(i, j int) bool {
	return (*p)[i].Val < (*p)[j].Val
}

func (p *PriorityQueue) Swap(i, j int) {
	(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(*ListNode))
}

func (p *PriorityQueue) Pop() any {
	last := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return last
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	/*
		放到小顶堆中
	*/
	d1 := &ListNode{-1, nil} // 虚拟头节点
	p1 := d1

	// 将所有链表的头节点加入堆
	p := &PriorityQueue{}
	heap.Init(p)

	for _, head := range lists {
		if head != nil {
			heap.Push(p, head)
		}
	}

	for p.Len() > 0 {
		node := heap.Pop(p).(*ListNode)
		p1.Next = node
		if node.Next != nil {
			heap.Push(p, node.Next)
		}
		p1 = p1.Next
	}

	return d1.Next
}
