package id3634

import "slices"

/*
 * @lc app=leetcode.cn id=3634 lang=golang
 *
 * [3634] Minimum Removals to Balance Array
 */

// @lc code=start
func minRemoval(nums []int, k int) int {
	slices.Sort(nums)
	left := 0
	maxKeep := 0
	for right, in := range nums {
		for ; in > nums[left]*k; left++ {
		}
		maxKeep = max(maxKeep, right-left+1)
	}
	return len(nums) - maxKeep
}

// @lc code=end
