package id50

/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	if n == -1 {
		return 1 / x
	}
	half := myPow(x, n/2)
	return half * half * myPow(x, n%2)
}

// @lc code=end
