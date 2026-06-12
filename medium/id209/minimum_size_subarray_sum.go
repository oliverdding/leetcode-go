package id209

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] Minimum Size Subarray Sum
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	sum := 0
	left := 0
	res := len(nums) + 1
	for right, in := range nums {
		sum += in
		for ; sum-nums[left] >= target; left++ {
			sum -= nums[left]
		}
		if sum >= target {
			res = min(res, right-left+1)
		}
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

// @lc code=end
