package id392

/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] Is Subsequence
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	sb := []byte(s)
	tb := []byte(t)
	si, ti := 0, 0
	for ; si < len(sb) && ti < len(tb); ti++ {
		if sb[si] == tb[ti] {
			si++
		}
	}
	if si == len(sb) {
		return true
	}
	return false
}

// @lc code=end
