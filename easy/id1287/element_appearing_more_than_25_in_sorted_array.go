package id1287

import "sort"

/*
 * @lc app=leetcode.cn id=1287 lang=golang
 *
 * [1287] Element Appearing More Than 25% In Sorted Array
 */

// @lc code=start
func findSpecialInteger(arr []int) int {
	n := len(arr)
	if sort.SearchInts(arr, arr[n/4]+1)-sort.SearchInts(arr, arr[n/4]) > n/4 {
		return arr[n/4]
	}
	if sort.SearchInts(arr, arr[n/2]+1)-sort.SearchInts(arr, arr[n/2]) > n/4 {
		return arr[n/2]
	}
	return arr[n*3/4]
}

// @lc code=end
