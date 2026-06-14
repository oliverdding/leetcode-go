package id2962

import "slices"

/*
 * @lc app=leetcode.cn id=2962 lang=golang
 *
 * [2962] Count Subarrays Where Max Element Appears at Least K Times
 */

// @lc code=start
func countSubarrays(nums []int, k int) int64 {
	me := slices.Max(nums)
	res, left, cnt := int64(0), 0, 0
	for _, in := range nums {
		if in == me {
			cnt++
		}
		for ; cnt >= k; left++ { // quit when condition not met
			if nums[left] == me {
				cnt--
			}
		}
		res += int64(left)
	}
	return res
}

// @lc code=end
