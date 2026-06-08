package id2779

import "slices"

/*
 * @lc app=leetcode.cn id=2779 lang=golang
 *
 * [2779] Maximum Beauty of an Array After Applying Operation
 */

// @lc code=start
func maximumBeauty(nums []int, k int) int {
	slices.Sort(nums)
	res := 0
	left := 0
	for right := 0; right < len(nums); right++ {
		for nums[right]-nums[left] > 2*k {
			left++
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
