package id3090

/*
 * @lc app=leetcode.cn id=3090 lang=golang
 *
 * [3090] Maximum Length Substring With Two Occurrences
 */

// @lc code=start
func maximumLengthSubstring(s string) int {
	apprs := [26]int{}
	res := 0
	left := 0
	for right, in := range s {
		apprs[in-'a']++
		for ; apprs[in-'a'] > 2; left++ {
			apprs[s[left]-'a']--
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
