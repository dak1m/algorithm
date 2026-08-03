package array

import "container/heap"

/*
https://leetcode.cn/problems/find-k-pairs-with-smallest-sums/

给定两个以 非递减顺序排列 的整数数组 nums1 和 nums2 , 以及一个整数 k 。

定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。

请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

示例 1:

输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
输出: [1,2],[1,4],[1,6]
解释: 返回序列中的前 3 对数：
     [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]
示例 2:

输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
输出: [1,1],[1,1]
解释: 返回序列中的前 2 对数：
     [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]
提示:

1 <= nums1.length, nums2.length <= 105
-109 <= nums1[i], nums2[i] <= 109
nums1 和 nums2 均为 升序排列
1 <= k <= 104
k <= nums1.length * nums2.length
题目来源：力扣 373. 查找和最小的 K 对数字。
*/

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	if len(nums1) == 0 {
		return nil
	}
	pl := &PriorityList{}
	heap.Init(pl)

	// 初始化列表
	for i := 0; i < len(nums1); i++ {
		heap.Push(pl, []int{nums1[i], nums2[0], 0})
	}

	result := make([][]int, 0)
	for pl.Len() > 0 && k > 0 {
		p := heap.Pop(pl).([]int)
		k--

		nextIndex := p[2] + 1
		if nextIndex < len(nums2) {
			heap.Push(pl, []int{p[0], nums2[nextIndex], nextIndex})
		}

		result = append(result, []int{p[0], p[1]})
	}

	return result
}

/*
PriorityList 实现一个最小堆
*/
type PriorityList [][]int

var _ heap.Interface = (*PriorityList)(nil)

func (p PriorityList) Len() int {
	return len(p)
}

func (p PriorityList) Less(i, j int) bool {
	return (p[i][0] + p[i][1]) < (p[j][0] + p[j][1])
}

func (p PriorityList) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityList) Push(x any) {
	*p = append(*p, x.([]int))
}

func (p *PriorityList) Pop() any {
	t := (*p)[len(*p)-1]
	*p = (*p)[0 : len(*p)-1]
	return t
}
