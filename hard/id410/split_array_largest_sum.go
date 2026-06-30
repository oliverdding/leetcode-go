package id410

import "slices"

/*
 * @lc app=leetcode.cn id=410 lang=golang
 *
 * [410] Split Array Largest Sum
 */

// @lc code=start
func splitArray(nums []int, k int) int {
	check := func(s int) bool {
		all, cnt := 0, 1
		for _, num := range nums {
			if all+num > s {
				all = 0
				cnt++
			}
			all += num
		}
		return cnt <= k
	}
	lo, hi := slices.Max(nums), 0
	for _, num := range nums {
		hi += num
	}
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
