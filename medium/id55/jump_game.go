package id55

/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	reach, maxJump := 0, nums[0]
	for reach != maxJump {
		if maxJump >= len(nums)-1 {
			return true
		}
		tmp := maxJump
		for i := reach + 1; i <= maxJump; i++ {
			tmp = max(tmp, i+nums[i])
		}
		reach = maxJump
		maxJump = tmp
	}
	return false
}

// @lc code=end
