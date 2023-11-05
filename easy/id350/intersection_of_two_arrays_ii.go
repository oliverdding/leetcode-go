package id350

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] Intersection of Two Arrays II
 */

// @lc code=start
func intersect(nums1 []int, nums2 []int) []int {
	occur := make([]int, 1001)
	for _, num1 := range nums1 {
		occur[num1]++
	}
	result := make([]int, 0)
	for _, num2 := range nums2 {
		if occur[num2]--; occur[num2] >= 0 {
			result = append(result, num2)
		}
	}
	return result
}

// @lc code=end
