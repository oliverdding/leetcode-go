package id35

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + ((hi - lo) >> 1)
		if nums[mid] > target {
			hi = mid
		} else if nums[mid] < target {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return lo
}

// @lc code=end
