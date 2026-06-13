package id713

/*
 * @lc app=leetcode.cn id=713 lang=golang
 *
 * [713] Subarray Product Less Than K
 */

// @lc code=start
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}
	res, left, prod := 0, 0, 1
	for right, in := range nums {
		prod *= in
		for ; prod >= k; left++ {
			prod /= nums[left]
		}
		res += right - left + 1 // [left~right, right] are all okay
	}
	return res
}

// @lc code=end
