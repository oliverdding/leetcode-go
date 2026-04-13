package id121

/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
func maxProfit(prices []int) int {
	minBuy := prices[0]
	maxEarn := 0
	for i := 1; i < len(prices); i++ {
		maxEarn = max(maxEarn, prices[i]-minBuy)
		minBuy = min(minBuy, prices[i])
	}
	return maxEarn
}

// @lc code=end
