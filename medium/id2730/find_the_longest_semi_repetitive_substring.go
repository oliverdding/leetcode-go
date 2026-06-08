package id2730

/*
 * @lc app=leetcode.cn id=2730 lang=golang
 *
 * [2730] Find the Longest Semi-Repetitive Substring
 */

// @lc code=start
func longestSemiRepetitiveSubstring(s string) int {
	res := 1
	cnt := 0
	left := 0
	for right := 1; right < len(s); right++ {
		if s[right] == s[right-1] {
			cnt++
		}
		for ; cnt > 1; left++ {
			if s[left] == s[left+1] {
				cnt--
			}
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
