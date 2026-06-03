package id1016

import (
	"math/bits"
	"strings"
)

/*
 * @lc app=leetcode.cn id=1016 lang=golang
 *
 * [1016] Binary String With Substrings Representing 1 To N
 */

// @lc code=start
func queryString(s string, n int) bool {
	if n == 1 {
		return strings.Contains(s, "1")
	}
	k := bits.Len(uint(n))
	m := len(s)
	if m < max(n+k-1<<(k-1), 1<<(k-2)+k-2) {
		return false
	}
	handle := func(k, low, high int) bool {
		v := 0
		mask := 1<<k - 1
		apprs := map[int]struct{}{}
		for right, in := range s {
			v = (v<<1 | int(in-'0')) & mask
			left := right - k + 1
			if left < 0 {
				continue
			}
			if low <= v && v <= high {
				apprs[v] = struct{}{}
			}
			if len(apprs) == (high - low + 1) {
				return true
			}
			if m-right-1 < high-low+1-len(apprs) {
				return false
			}
		}
		return false
	}
	return handle(k-1, n/2+1, 1<<(k-1)-1) && handle(k, 1<<(k-1), n)
}

// @lc code=end
