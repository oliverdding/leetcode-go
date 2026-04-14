package id15

import "slices"

/*
 * @lc app=leetcode.cn id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
func threeSum(nums []int) [][]int {
	slices.SortFunc(nums, func(a int, b int) int {
		return a - b
	})
	res := [][]int{}
	first := 0
	for first <= len(nums)-3 {
		if nums[first] > 0 {
			break
		}
		if nums[len(nums)-1] < 0 {
			break
		}
		second, third := first+1, len(nums)-1
		for second < third {
			s := nums[first] + nums[second] + nums[third]
			if s < 0 {
				second++
			} else if s > 0 {
				third--
			} else {
				res = append(res, []int{nums[first], nums[second], nums[third]})
				second++
				for second < third && nums[second] == nums[second-1] {
					second++
				}
				third--
				for second < third && nums[third] == nums[third+1] {
					third--
				}
			}
		}
		first++
		for first <= len(nums)-3 && nums[first] == nums[first-1] {
			first++
		}
	}
	return res
}

// @lc code=end
