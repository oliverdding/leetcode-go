package id34

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] Find First and Last Position of Element in Sorted Array
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	start, ok := lowerBound(nums, target)
	if !ok {
		return []int{-1, -1}
	}
	end, _ := lowerBound(nums, target+1)
	return []int{start, end - 1}
}

func lowerBound(nums []int, target int) (int, bool) {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo, lo < len(nums) && nums[lo] == target
}

// @lc code=end
