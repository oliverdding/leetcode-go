package id852

/*
 * @lc app=leetcode.cn id=852 lang=golang
 *
 * [852] Peak Index in a Mountain Array
 */

// @lc code=start
func peakIndexInMountainArray(arr []int) int {
	// [,)
	left, right := 1, len(arr)-1
	for left < right {
		mid := left + (right-left)>>1
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}

// @lc code=end
