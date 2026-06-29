package id2982

import (
	"slices"
)

/*
 * @lc app=leetcode.cn id=2982 lang=golang
 *
 * [2982] Find Longest Special Substring That Occurs Thrice II
 */

// @lc code=start
func maximumLength(s string) int {
	chunkCnt := make(map[byte][]int)

	cnt := 0
	for i := range s {
		cnt++
		if i+1 == len(s) || s[i] != s[i+1] {
			chunkCnt[s[i]] = append(chunkCnt[s[i]], cnt)
			cnt = 0
		}
	}

	res := 0
	for _, arr := range chunkCnt {
		check := func(g int) bool {
			sum := 0
			for _, y := range arr {
				sum += max(y-g+1, 0)
			}
			return sum >= 3
		}
		lo, hi := 1, slices.Max(arr)+1
		for lo < hi {
			mid := lo + (hi-lo)>>1
			if check(mid) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
		res = max(res, lo-1)
	}
	if res == 0 {
		return -1
	}
	return res
}

// @lc code=end
