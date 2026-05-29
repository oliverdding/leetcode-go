package id3652

import "math"

/*
 * @lc app=leetcode.cn id=3652 lang=golang
 *
 * [3652] Best Time to Buy and Sell Stock using Strategy
 */

// @lc code=start
func maxProfit(prices []int, strategy []int, k int) int64 {
	res := int64(0)
	s := len(prices)
	for i := 0; i < s; i++ {
		res += int64(prices[i]) * int64(strategy[i])
	}
	origin, modified := int64(0), int64(0)
	var maxChange int64 = math.MinInt64
	for right := 0; right < s; right++ {
		origin += int64(prices[right]) * int64(strategy[right])
		if right >= k/2 {
			modified += int64(prices[right])
		}
		left := right - k + 1
		if left < 0 {
			continue
		}
		maxChange = max(maxChange, res+modified-origin)
		origin -= int64(prices[left]) * int64(strategy[left])
		modified -= int64(prices[left+k/2])
	}
	res = max(res, maxChange)
	return int64(res)
}

// @lc code=end
