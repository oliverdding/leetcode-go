package id56

import (
	"cmp"
	"slices"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	res := make([][]int, 0)
	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})
	res = append(res, intervals[0])
	resPtr := 0
	for i := 1; i < len(intervals); i++ {
		if res[resPtr][1] >= intervals[i][0] {
			res[resPtr][1] = max(res[resPtr][1], intervals[i][1])
		} else {
			res = append(res, intervals[i])
			resPtr++
		}
	}
	return res
}

// @lc code=end
