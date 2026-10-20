package list

/*
给你一个字符串 s，找到 s 中最长的 回文 子串。

示例 1：

输入：s = "babad"
输出："bab"
解释："aba" 同样是符合题意的答案。
示例 2：

输入：s = "cbbd"
输出："bb"
提示：

1 <= s.length <= 1000
s 仅由数字和英文字母组成
题目来源：力扣 5. 最长回文子串。
https://leetcode.cn/problems/longest-palindromic-substring/
*/

func longestPalindrome(s string) string {
	/*
		整体思路：从中心点展开，取到最长的子串
		注意偶数和奇数，中心点分别为两个和一个
		左右指针从中心向两端扩展
	*/
	res := ""
	for i := 0; i < len(s); i++ {
		p1 := palindrome(s, i, i)
		if len(res) < len(p1) {
			res = p1
		}
		p2 := palindrome(s, i, i+1)
		if len(res) < len(p2) {
			res = p2
		}
	}

	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}

	return s[l+1 : r]
}
