package id2563

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2563 lang=golang
 *
 * [2563] Count the Number of Fair Pairs
 */

// @lc code=start
func countFairPairs(nums []int, lower int, upper int) int64 {
	n := len(nums)
	slices.Sort(nums)
	res := int64(0)
	for i := 0; i < n-1; i++ {
		res += int64(sort.SearchInts(nums[i+1:], upper-nums[i]+1)) - int64(sort.SearchInts(nums[i+1:], lower-nums[i]))
	}
	return res
}

// @lc code=end
