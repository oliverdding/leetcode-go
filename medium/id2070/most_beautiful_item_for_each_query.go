package id2070

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2070 lang=golang
 *
 * [2070] Most Beautiful Item for Each Query
 */

// @lc code=start
func maximumBeauty(items [][]int, queries []int) []int {
	n := len(items)
	slices.SortFunc(items, func(a []int, b []int) int {
		return a[0] - b[0]
	})
	prefixMax := make([]int, n)
	prefixMax[0] = items[0][1]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], items[i][1])
	}
	res := make([]int, 0, n)
	for _, query := range queries {
		idx := sort.Search(n, func(i int) bool {
			return items[i][0] > query
		}) - 1
		if idx >= 0 {
			res = append(res, prefixMax[idx])
		} else {
			res = append(res, 0)
		}
	}
	return res
}

// @lc code=end
