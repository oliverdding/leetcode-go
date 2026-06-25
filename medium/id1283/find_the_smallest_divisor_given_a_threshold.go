package id1283

import "slices"

/*
 * @lc app=leetcode.cn id=1283 lang=golang
 *
 * [1283] Find the Smallest Divisor Given a Threshold
 */

// @lc code=start
func smallestDivisor(nums []int, threshold int) int {
	check := func(div int) bool {
		res := 0
		for _, num := range nums {
			res += (num + div - 1) / div
			if res > threshold {
				return false
			}
		}
		return true
	}
	lo, hi := 1, slices.Max(nums)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end
