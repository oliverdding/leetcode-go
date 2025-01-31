package id455

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=455 lang=golang
 *
 * [455] Assign Cookies
 */

// @lc code=start
func findContentChildren(g []int, s []int) int {
	slices.Sort(g)
	slices.Sort(s)

	i := 0
	for _, size := range s {
		if i < len(g) && g[i] <= size {
			i++
		}
	}

	return i
}

// @lc code=end
