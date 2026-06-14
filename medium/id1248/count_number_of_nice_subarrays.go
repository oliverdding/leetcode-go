package id1248

/*
 * @lc app=leetcode.cn id=1248 lang=golang
 *
 * [1248] Count Number of Nice Subarrays
 */

// @lc code=start
func numberOfSubarrays(nums []int, k int) int {
	return solve(nums, k) - solve(nums, k-1)
}

func solve(nums []int, k int) int {
	res, sum, left := 0, 0, 0
	for right, in := range nums {
		sum += in & 1
		for ; sum > k; left++ { // quit when condition met
			sum -= nums[left] & 1
		}
		res += right - left + 1
	}
	return res
}

// @lc code=end
