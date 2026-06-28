package id2226

import "slices"

/*
 * @lc app=leetcode.cn id=2226 lang=golang
 *
 * [2226] Maximum Candies Allocated to K Children
 */

// @lc code=start
func maximumCandies(candies []int, k int64) int {
	check := func(m int) bool {
		able := 0
		for _, candy := range candies {
			able += candy / m
			if able >= int(k) {
				return true
			}
		}
		return false
	}
	lo, hi := 1, slices.Max(candies)/((int(k)-1)/len(candies)+1)+1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo - 1
}

// @lc code=end
