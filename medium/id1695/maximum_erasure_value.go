package id1695

/*
 * @lc app=leetcode.cn id=1695 lang=golang
 *
 * [1695] Maximum Erasure Value
 */

// @lc code=start
func maximumUniqueSubarray(nums []int) int {
	res := 0
	left := 0
	sum := 0
	apprs := [10001]bool{}
	for _, in := range nums {
		for ; apprs[in]; left++ {
			out := nums[left]
			apprs[out] = false
			sum -= out
		}
		apprs[in] = true
		sum += in
		res = max(res, sum)
	}
	return res
}

// @lc code=end
