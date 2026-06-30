package id875

import (
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=875 lang=golang
 *
 * [875] Koko Eating Bananas
 */

// @lc code=start
func minEatingSpeed(piles []int, h int) int {
	n := len(piles)
	q := h / n
	lo, hi := max((slices.Min(piles)+q-1)/q-1, 1), (slices.Max(piles)+q-1)/q
	return lo + sort.Search(hi-lo, func(k int) bool {
		k = lo + k
		cost := 0
		for _, pile := range piles {
			cost += (pile + k - 1) / k
			if cost > h {
				return false
			}
		}
		return true
	})
}

// @lc code=end
