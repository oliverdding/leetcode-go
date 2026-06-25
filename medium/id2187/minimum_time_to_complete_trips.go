package id2187

import "slices"

/*
 * @lc app=leetcode.cn id=2187 lang=golang
 *
 * [2187] Minimum Time to Complete Trips
 */

// @lc code=start
func minimumTime(time []int, totalTrips int) int64 {
	check := func(t int64) bool {
		trips := 0
		for i := range time {
			trips += int(t / int64(time[i]))
			if trips >= totalTrips {
				return true
			}
		}
		return false
	}
	lo, hi := int64(slices.Min(time)-1), int64(slices.Min(time))*int64(totalTrips)
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
