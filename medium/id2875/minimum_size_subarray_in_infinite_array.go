package id2875

import "math"

/*
 * @lc app=leetcode.cn id=2875 lang=golang
 *
 * [2875] Minimum Size Subarray in Infinite Array
 */

// @lc code=start
func minSizeSubarray(nums []int, target int) int {
	total := 0
	for _, in := range nums {
		total += in
	}
	left := 0
	sum := 0
	res := math.MaxInt
	for right := 0; right < len(nums)*2; right++ {
		sum += nums[right%len(nums)]
		for ; sum > target%total; left++ {
			sum -= nums[left%len(nums)]
		}
		if sum == target%total {
			res = min(res, right-left+1)
		}
	}
	if res == math.MaxInt {
		return -1
	}
	return res + len(nums)*(target/total)
}

// @lc code=end
