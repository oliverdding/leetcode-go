package id1838

import "slices"

/*
 * @lc app=leetcode.cn id=1838 lang=golang
 *
 * [1838] Frequency of the Most Frequent Element
 */

// @lc code=start
func maxFrequency(nums []int, k int) int {
	slices.Sort(nums)
	left := 0
	res := 0
	sum := 0
	for right, in := range nums {
		sum += in
		for ; in*(right-left+1)-sum > k; left++ {
			sum -= nums[left]
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
