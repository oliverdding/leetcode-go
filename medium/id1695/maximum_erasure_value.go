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
	apprs := [10001]int{}
	for _, in := range nums {
		apprs[in]++
		sum += in
		for ; apprs[in] != 1; left++ {
			out := nums[left]
			apprs[out]--
			sum -= out
		}
		res = max(res, sum)
	}
	return res
}

// @lc code=end
