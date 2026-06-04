package id2653

/*
 * @lc app=leetcode.cn id=2653 lang=golang
 *
 * [2653] Sliding Subarray Beauty
 */

// @lc code=start
func getSubarrayBeauty(nums []int, k int, x int) []int {
	n := len(nums)
	res := make([]int, n-k+1)
	cnt := [101]int{}

	for right, in := range nums {
		cnt[in+50]++
		left := right - k + 1
		if left < 0 {
			continue
		}
		sum := 0
		for i := 0; i < 50; i++ {
			sum += cnt[i]
			if sum >= x {
				res[right+1-k] = i - 50
				break
			}
		}
		cnt[nums[left]+50]--
	}
	return res
}

// @lc code=end
