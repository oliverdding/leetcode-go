package id69

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	lo, hi := 1, x+1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if mid > x/mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}

// @lc code=end
