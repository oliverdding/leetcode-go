package id1385

import "slices"

/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] Find the Distance Value Between Two Arrays
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	slices.Sort(arr2)
	for _, num := range arr1 {
		idx := lowerBound(arr2, num-d)
		if idx == len(arr2) || arr2[idx] > num+d {
			res++
		}
	}
	return res
}

func lowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end
