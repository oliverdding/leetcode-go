package id658

import "sort"

/*
 * @lc app=leetcode.cn id=658 lang=golang
 *
 * [658] Find K Closest Elements
 */

// @lc code=start
func findClosestElements(arr []int, k int, x int) []int {
	idx := sort.SearchInts(arr, x)
	left, right := idx, idx
	if idx == len(arr) || (idx > 0 && abs(arr[idx-1]-x) <= abs(arr[idx]-x)) {
		left, right = idx-1, idx-1
	}
	for right-left+1 < k {
		if left == 0 {
			right++
			continue
		}
		if right == len(arr)-1 {
			left--
			continue
		}
		if abs(arr[left-1]-x) <= abs(arr[right+1]-x) {
			left--
		} else {
			right++
		}
	}
	return arr[left : right+1]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// @lc code=end
