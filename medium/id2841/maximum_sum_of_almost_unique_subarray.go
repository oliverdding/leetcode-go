package id2841

/*
 * @lc app=leetcode.cn id=2841 lang=golang
 *
 * [2841] Maximum Sum of Almost Unique Subarray
 */

// @lc code=start
func maxSum(nums []int, m int, k int) int64 {
	cntMap := make(map[int]int)
	res := int64(0)
	sum := int64(0)
	for right, in := range nums {
		cntMap[in]++
		sum += int64(in)

		left := right - k + 1
		if left < 0 {
			continue
		}
		if len(cntMap) >= m {
			res = max(res, sum)
		}

		out := nums[left]
		cntMap[out] = cntMap[out] - 1
		if cntMap[out] == 0 {
			delete(cntMap, out)
		}
		sum -= int64(out)
	}
	return res
}

// @lc code=end
