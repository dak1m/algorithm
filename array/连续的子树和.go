package array

/*
给你一个整数数组 nums 和一个整数 k ，如果 nums 有一个 好的子数组 返回 true ，否则返回 false：

一个 好的子数组 是：

长度 至少为 2 ，且
子数组元素总和为 k 的倍数。
注意：

子数组 是数组中 连续 的部分。
如果存在一个整数 n ，令整数 x 符合 x = n * k ，则称 x 是 k 的一个倍数。0 始终 视为 k 的一个倍数。


示例 1：

输入：nums = [23,2,4,6,7], k = 6
输出：true
解释：[2,4] 是一个大小为 2 的子数组，并且和为 6 。
示例 2：

输入：nums = [23,2,6,4,7], k = 6
输出：true
解释：[23, 2, 6, 4, 7] 是大小为 5 的子数组，并且和为 42 。
42 是 6 的倍数，因为 42 = 7 * 6 且 7 是一个整数。
示例 3：

输入：nums = [23,2,6,4,7], k = 13
输出：false


提示：

1 <= nums.length <= 105
0 <= nums[i] <= 109
0 <= sum(nums[i]) <= 231 - 1
1 <= k <= 231 - 1
*/

/*
此方法会超时
*/
func checkSubarraySum(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	for i := 2; i < len(preSum); i++ {
		for j := 0; j < i-1; j++ {
			if preSum[i]%k == 0 {
				return true
			}
			if (preSum[i]-preSum[j])%k == 0 {
				return true
			}
		}
	}

	return false
}

/*
(preSum[i] - preSum[j]) % k == 0

	=> preSum[i] % k = preSum[j] % k
*/
func checkSubarraySumByHashMap(nums []int, k int) bool {
	preSum := make([]int, len(nums)+1)
	for i := 1; i < len(preSum); i++ {
		preSum[i] = preSum[i-1] + nums[i-1]
	}

	preSumMap := make(map[int]int)
	for i := 0; i < len(preSum); i++ {
		p := preSum[i] % k
		idx, ok := preSumMap[p]
		if !ok {
			preSumMap[p] = i
		} else {
			if i-idx >= 2 {
				return true
			}
		}
	}

	//for i := 0; i < len(preSum); i++ {
	//	mod := preSum[i] % k
	//	if index, exists := preSumMap[mod]; exists {
	//		if i-index >= 2 {
	//			return true
	//		}
	//	} else {
	//		preSumMap[mod] = i
	//	}
	//}

	return false
}
