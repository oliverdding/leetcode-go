package id11

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
func maxArea(height []int) int {
	maximum := 0
	lo, hi := 0, len(height)-1
	for lo != hi {
		maximum = max(maximum, (hi-lo)*min(height[lo], height[hi]))
		if height[lo] >= height[hi] {
			hi--
		} else {
			lo++
		}
	}
	return maximum
}

// @lc code=end
