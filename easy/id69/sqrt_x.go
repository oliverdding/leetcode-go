package id69

/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) int {
	low, high := 1, x
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid > x {
			high = mid - 1
		} else if mid*mid < x {
			low = mid + 1
		} else {
			return mid
		}
	}
	return high
}

// @lc code=end
