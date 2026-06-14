package id3258

/*
 * @lc app=leetcode.cn id=3258 lang=golang
 *
 * [3258] Count Substrings That Satisfy K-Constraint I
 */

// @lc code=start
func countKConstraintSubstrings(s string, k int) int {
	res, cnt, left := 0, [2]int{}, 0
	for right, in := range s {
		cnt[in&1]++
		for ; cnt[0] > k && cnt[1] > k; left++ { // exit when satisfied
			cnt[s[left]&1]--
		}
		res += right - left + 1
	}
	return res
}

// @lc code=end
