package id3296

import "slices"

/*
 * @lc app=leetcode.cn id=3296 lang=golang
 *
 * [3296] Minimum Number of Seconds to Make Mountain Height Zero
 */

// @lc code=start
func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	check := func(t int) bool {
		done := 0
		for _, workerTime := range workerTimes {
			done += (sqrt(1+8*(t/workerTime)) - 1) / 2
			if done >= int(mountainHeight) {
				return true
			}
		}
		return false
	}
	h := (mountainHeight-1)/len(workerTimes) + 1
	lo, hi := 0, slices.Max(workerTimes)*(1+h)*h/2
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return int64(lo)
}

func sqrt(x int) int {
	lo, hi := 1, x+1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if mid > x/mid {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo - 1
}

// @lc code=end
