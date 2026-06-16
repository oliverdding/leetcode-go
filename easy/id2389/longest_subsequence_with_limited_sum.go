package id2389

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=2389 lang=golang
 *
 * [2389] Longest Subsequence With Limited Sum
 */

// @lc code=start
func answerQueries(nums []int, queries []int) []int {
	slices.Sort(nums)
	prefixs := make([]int, len(nums))
	for i, num := range nums {
		if i == 0 {
			prefixs[0] = num
			continue
		}
		prefixs[i] = prefixs[i-1] + num
	}
	res := make([]int, 0, len(queries))
	for _, query := range queries {
		res = append(res, sort.SearchInts(prefixs, query+1))
	}
	return res
}

// @lc code=end
