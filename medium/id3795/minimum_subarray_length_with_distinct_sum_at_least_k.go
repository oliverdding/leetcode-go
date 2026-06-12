package id3795

/*
 * @lc app=leetcode.cn id=3795 lang=golang
 *
 * [3795] Minimum Subarray Length With Distinct Sum At Least K
 */

// @lc code=start
func minLength(nums []int, k int) int {
	res, left, sum := len(nums)+1, 0, 0
	apprs := map[int]int{}
	for right, in := range nums {
		apprs[in]++
		if apprs[in] == 1 {
			sum += in
		}
		for ; sum >= k; left++ {
			apprs[nums[left]]--
			if apprs[nums[left]] == 0 {
				sum -= nums[left]
				delete(apprs, nums[left])
			}
			res = min(res, right-left+1)
		}
	}
	if res == len(nums)+1 {
		return -1
	}
	return res
}

// @lc code=end
