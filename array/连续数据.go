package array

/*
思路：把1看做-1，题目就变为了和为0的最长子数组
使得 preSum[i] - preSum[j] == 0，其中 i > j, i-j 最大
*/
func findMaxLength(nums []int) int {
	preSum := make([]int, len(nums)+1)
	res := 0
	for i := 1; i < len(preSum); i++ {
		if nums[i-1] == 0 {
			preSum[i] = preSum[i-1] - 1
		} else {
			preSum[i] = preSum[i-1] + 1
		}
	}
	preSumMap := make(map[int]int, len(nums))
	for i, v := range preSum {
		if val, ok := preSumMap[v]; !ok {
			preSumMap[v] = i
		} else {
			res = maxInt(res, i-val)
		}
	}

	return res
}

func findMaxLengthAdvanced(nums []int) int {
	preSum := 0
	res := 0
	preSumMap := map[int]int{
		0: 0,
	}
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			preSum--
		} else {
			preSum++
		}
		if val, ok := preSumMap[preSum]; ok {
			res = maxInt(res, i-val)
		} else {
			preSumMap[preSum] = i
		}
	}
	return res
}

//
//func findMaxLengthAdvanced(nums []int) int {
//	// 使用变量代替数组，空间复杂度优化
//	preSum := 0
//	res := 0
//	// preSumMap 存储前缀和第一次出现的位置
//	preSumMap := make(map[int]int)
//	// 初始化：前缀和为0时，位置为0（表示空数组）
//	preSumMap[0] = 0
//
//	for i := 0; i < len(nums); i++ {
//		// 更新当前前缀和：0看作-1，1看作+1
//		if nums[i] == 0 {
//			preSum--
//		} else {
//			preSum++
//		}
//
//		// 如果当前前缀和之前出现过，说明中间这段子数组的和为0
//		if val, ok := preSumMap[preSum]; ok {
//			res = maxInt(res, i+1-val)
//		} else {
//			// 记录这个前缀和第一次出现的位置
//			preSumMap[preSum] = i + 1
//		}
//	}
//
//	return res
//}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
