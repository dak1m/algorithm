package array

/*
查询分数在 [80, 90] 之间的同学有多少人
*/
func scoreRangeQuery(scores []int, lower int, upper int) int {
	// 试卷满分 100 分
	count := make([]int, 100+1)

	// 记录每个分数有几个同学
	for _, score := range scores {
		count[score]++
	}
	// 构造前缀和数组
	for i := 1; i < len(count); i++ {
		count[i] = count[i] + count[i-1]
	}

	// 查询分数在 [lower, upper] 之间的同学有多少人
	result := count[upper] - count[lower-1]
	return result
}
