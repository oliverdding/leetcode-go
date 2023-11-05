package id367

/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] Valid Perfect Square
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	low, high := 1, num
	for low <= high {
		mid := low + (high-low)/2
		if num/mid > mid {
			low = mid + 1
		} else if num/mid < mid {
			high = mid - 1
		} else if num%mid == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

// @lc code=end
