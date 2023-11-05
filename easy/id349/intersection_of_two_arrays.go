package id349

/*
 * @lc app=leetcode.cn id=349 lang=golang
 *
 * [349] Intersection of Two Arrays
 */

// @lc code=start
func intersection(nums1 []int, nums2 []int) []int {
	occur1 := make([]bool, 1001)
	for _, num1 := range nums1 {
		occur1[num1] = true
	}
	occur2 := make([]bool, 1001)
	for _, num2 := range nums2 {
		occur2[num2] = true
	}
	result := make([]int, 0)
	for idx := 0; idx <= 1000; idx++ {
		if occur1[idx] && occur2[idx] {
			result = append(result, idx)
		}
	}
	return result
}

// @lc code=end
