package id3652

/*
 * @lc app=leetcode.cn id=3652 lang=golang
 *
 * [3652] Best Time to Buy and Sell Stock using Strategy
 */

// @lc code=start
func maxProfit(prices []int, strategy []int, k int) int64 {
	total := int64(0)
	diff := int64(0)
	maxDiff := int64(0)
	for i := 0; i < len(prices); i++ {
		total += int64(prices[i]) * int64(strategy[i])

		diff += int64(prices[i]) * (1 - int64(strategy[i])) // enter right part, strategy -> 1
		left := i - k + 1
		if left < 0 { // window not fill yet
			if left+k/2 >= 0 { // right element (left+k/2) enter left part, 1 -> 0
				diff += int64(-prices[left+k/2])
			}
			continue
		}
		maxDiff = max(maxDiff, diff)
		// left element leave window, 0 -> strategy
		// and right element (left+k/2) enter left part, 1 -> 0
		diff += int64(-prices[left+k/2]) + int64(prices[left])*int64(strategy[left])
	}
	return total + maxDiff
}

// @lc code=end
