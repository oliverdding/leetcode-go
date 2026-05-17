package id2090

/*
 * @lc app=leetcode.cn id=2090 lang=golang
 *
 * [2090] K Radius Subarray Averages
 */

// @lc code=start
func getAverages(nums []int, k int) []int {
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	sum := 0
	for right, in := range nums {
		sum += in
		left := right - 2*k
		if left < 0 {
			continue
		}
		res[right-k] = sum / (2*k + 1)
		sum -= nums[left]
	}
	return res
}

// @lc code=end
