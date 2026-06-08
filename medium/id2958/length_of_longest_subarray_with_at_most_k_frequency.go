package id2958

/*
 * @lc app=leetcode.cn id=2958 lang=golang
 *
 * [2958] Length of Longest Subarray With at Most K Frequency
 */

// @lc code=start
func maxSubarrayLength(nums []int, k int) int {
	res := 0
	left := 0
	apprs := make(map[int]int)
	for right, in := range nums {
		apprs[in]++
		for ; apprs[in] > k; left++ {
			out := nums[left]
			apprs[out]--
			if apprs[out] == 0 {
				delete(apprs, out)
			}
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
