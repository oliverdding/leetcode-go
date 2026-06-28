package id3824

import (
	"math"
	"slices"
)

/*
 * @lc app=leetcode.cn id=3824 lang=golang
 *
 * [3824] Minimum K to Reduce Array Within Limit
 */

// @lc code=start
func minimumK(nums []int) int {
	n := len(nums)
	check := func(k int) bool {
		nonPositive := n
		for _, num := range nums {
			nonPositive += (num - 1) / k
			if nonPositive > k*k {
				return false
			}
		}
		return true
	}
	lo, hi := int(math.Ceil(math.Sqrt(float64(n)))), slices.Max(nums)
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
