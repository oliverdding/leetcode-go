package id441

/*
 * @lc app=leetcode.cn id=441 lang=golang
 *
 * [441] Arranging Coins
 */

// @lc code=start
func arrangeCoins(n int) int {
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2
		if coins := (1 + mid) * mid / 2; coins < n {
			low = mid + 1
		} else if coins > n {
			high = mid - 1
		} else {
			return mid
		}
	}
	return high
}

// @lc code=end
