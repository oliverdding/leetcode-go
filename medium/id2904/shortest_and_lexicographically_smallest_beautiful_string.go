package id2904

/*
 * @lc app=leetcode.cn id=2904 lang=golang
 *
 * [2904] Shortest and Lexicographically Smallest Beautiful String
 */

// @lc code=start
func shortestBeautifulSubstring(s string, k int) string {
	size, left, cnt1 := len(s)+1, 0, 0
	minArr := ""
	for right, in := range s {
		cnt1 += int(in) & 1
		for ; cnt1 > k; left++ {
			cnt1 -= int(s[left]) & 1
		}
		if cnt1 == k {
			for ; s[left] == '0'; left++ {
			}
			if size == right-left+1 {
				minArr = min(minArr, s[left:right+1])
			} else if size > right-left+1 {
				size = right - left + 1
				minArr = s[left : right+1]
			}
		}
	}
	return minArr
}

// @lc code=end
