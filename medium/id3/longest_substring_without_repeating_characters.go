package id3

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	res := 0
	apprs := [95]bool{}
	for left, right := 0, 0; right < len(s); right++ {
		for apprs[s[right]-' '] {
			apprs[s[left]-' '] = false
			left++
		}
		res = max(res, right-left+1)
		apprs[s[right]-' '] = true
	}

	return res
}

// @lc code=end
