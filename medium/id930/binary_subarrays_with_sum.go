package id930

/*
 * @lc app=leetcode.cn id=930 lang=golang
 *
 * [930] Binary Subarrays With Sum
 */

// @lc code=start
func numSubarraysWithSum(nums []int, goal int) int {
	return solve(nums, goal) - solve(nums, goal+1)
}

func solve(nums []int, goal int) int {
	res, sum, left := 0, 0, 0
	for right, in := range nums {
		sum += in
		for ; sum >= goal && left <= right; left++ { // quit when condition not met
			sum -= nums[left]
		}
		res += left
	}
	return res
}

// @lc code=end
