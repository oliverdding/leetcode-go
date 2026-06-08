package id1004

/*
 * @lc app=leetcode.cn id=1004 lang=golang
 *
 * [1004] Max Consecutive Ones III
 */

// @lc code=start
func longestOnes(nums []int, k int) int {
	res := 0
	cnt1 := 0
	left := 0
	for right, in := range nums {
		cnt1 += in
		for ; right-left+1-cnt1 > k; left++ {
			cnt1 -= nums[left]
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
