package id643

import "math"

/*
 * @lc app=leetcode.cn id=643 lang=golang
 *
 * [643] Maximum Average Subarray I
 */

// @lc code=start
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	maxSum := math.MinInt
	for right, in := range nums {
		sum += in
		left := right - k + 1
		if left < 0 {
			continue
		}
		maxSum = max(maxSum, sum)
		sum -= nums[left]
	}
	return float64(maxSum) / float64(k)
}

// @lc code=end
