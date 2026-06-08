package id1208

/*
 * @lc app=leetcode.cn id=1208 lang=golang
 *
 * [1208] Get Equal Substrings Within Budget
 */

// @lc code=start
func equalSubstring(s string, t string, maxCost int) int {
	n := len(s)
	res := 0
	cost := 0
	left := 0
	for right := 0; right < n; right++ {
		cost += subtractWithABS(s[right], t[right])
		for ; cost > maxCost; left++ {
			cost -= subtractWithABS(s[left], t[left])
		}
		res = max(res, right-left+1)
	}
	return res
}

func subtractWithABS(a, b byte) int {
	if a > b {
		return int(a - b)
	}
	return int(b - a)
}

// @lc code=end
