package id390

/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] Elimination Game
 */

// @lc code=start
func lastRemaining(n int) int {
	start, step := 1, 1
	for n > 1 {
		start += (n - 1 - n&1) * step
		step *= -2
		n >>= 1
	}
	return start
}

// @lc code=end
