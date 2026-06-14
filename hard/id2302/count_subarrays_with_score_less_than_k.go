package id2302

/*
 * @lc app=leetcode.cn id=2302 lang=golang
 *
 * [2302] Count Subarrays With Score Less Than K
 */

// @lc code=start
func countSubarrays(nums []int, k int64) int64 {
	res, sum, left := 0, int64(0), 0
	for right, in := range nums {
		sum += int64(in)
		for ; sum*int64(right-left+1) >= k; left++ {
			sum -= int64(nums[left])
		}
		res += right - left + 1
	}
	return int64(res)
}

// @lc code=end
