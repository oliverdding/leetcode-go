package id209

import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := math.MaxInt
	left, right, sum := 0, 0, 0
	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			res = min(res, right - left + 1)
			sum -= nums[left]
			left++
		}
		right++
	}
	if res ==  math.MaxInt {
		return 0
	}
	return res
}
// @lc code=end

