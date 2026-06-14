package id3325

/*
 * @lc app=leetcode.cn id=3325 lang=golang
 *
 * [3325] Count Substrings With K-Frequency Characters I
 */

// @lc code=start
func numberOfSubstrings(s string, k int) int {
	cnt := [27]int{}
	res, left := 0, 0
	for right, in := range s {
		cnt[in&31]++
		for ; cnt[in&31] >= k; left++ { // quit when condition not met
			res += len(s) - right
			cnt[s[left]&31]--
		}
	}
	return res
}

// @lc code=end
