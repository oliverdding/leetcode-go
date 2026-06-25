package id1011

import "slices"

/*
 * @lc app=leetcode.cn id=1011 lang=golang
 *
 * [1011] Capacity To Ship Packages Within D Days
 */

// @lc code=start
func shipWithinDays(weights []int, days int) int {
	check := func(cap int) bool {
		need := 0
		carry := 0
		for _, weight := range weights {
			if carry+weight > cap {
				need++ // launch ship
				carry = 0
				if need > days {
					return false
				}
			}
			carry += weight
		}
		if carry > 0 { // last round trip
			need++
		}
		return need <= days
	}
	n := len(weights)
	lo := slices.Max(weights)
	hi := (((lo*n+days-1)/days + lo - 1) / lo) * lo
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return hi
}

// @lc code=end
