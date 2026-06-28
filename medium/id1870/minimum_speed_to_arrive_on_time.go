package id1870

import (
	"math"
	"slices"
	"sort"
)

/*
 * @lc app=leetcode.cn id=1870 lang=golang
 *
 * [1870] Minimum Speed to Arrive on Time
 */

// @lc code=start
func minSpeedOnTime(dist []int, hour float64) int {
	n := len(dist)
	if hour <= float64(n-1) {
		return -1
	}
	lo, hi := 1, max(slices.Max(dist), (100*dist[n-1]-1)/(int(math.Round(hour*100))-100*(n-1))+1)
	return lo + sort.Search(hi-lo, func(k int) bool {
		k += lo
		cost := n - 1
		for i := 0; i < n-1; i++ {
			cost += (dist[i] - 1) / k
			if cost > int(hour) {
				return false
			}
		}
		return 100*(cost*k+dist[n-1]) <= int(math.Round(hour*100))*k
	})
}

// @lc code=end
