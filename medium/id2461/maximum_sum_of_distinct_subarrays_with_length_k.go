package id2461

/*
 * @lc app=leetcode.cn id=2461 lang=golang
 *
 * [2461] Maximum Sum of Distinct Subarrays With Length K
 */

// @lc code=start
func maximumSubarraySum(nums []int, k int) int64 {
	appr := map[int]int{}
	sum := int64(0)
	res := int64(0)
	for right, in := range nums {
		appr[in]++
		sum += int64(in)
		left := right - k + 1
		if left < 0 {
			continue
		}

		if len(appr) == k {
			res = max(res, sum)
		}

		out := nums[left]
		appr[out]--
		if appr[out] == 0 {
			delete(appr, out)
		}
		sum -= int64(out)
	}
	return int64(res)
}

// @lc code=end
