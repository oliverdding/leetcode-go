package id2200

/*
 * @lc app=leetcode.cn id=2200 lang=golang
 *
 * [2200] Find All K-Distant Indices in an Array
 */

// @lc code=start
func findKDistantIndices(nums []int, key int, k int) []int {
	s := len(nums)
	res := []int{}
	mark := 0
	for i := 0; i < s+k; i++ {
		if i < s && nums[i] == key {
			mark++
		}
		left := i - 2*k
		if left < 0 {
			if mark > 0 && left+k >= 0 {
				res = append(res, left+k)
			}
			continue
		}
		if mark > 0 {
			res = append(res, left+k)
		}
		if nums[left] == key {
			mark--
		}
	}
	return res
}

// @lc code=end
