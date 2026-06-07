package id1493

/*
 * @lc app=leetcode.cn id=1493 lang=golang
 *
 * [1493] Longest Subarray of 1's After Deleting One Element
 */

// @lc code=start
func longestSubarray(nums []int) int {
	cntZero := 0
	left := 0
	res := 0
	for right, in := range nums {
		cntZero += 1 - in
		for ; cntZero == 2; left++ {
			cntZero -= 1 - nums[left]
		}
		res = max(res, right-left)
	}
	return res
}

// @lc code=end
