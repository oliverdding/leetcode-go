package id3143

import "math/bits"

/*
 * @lc app=leetcode.cn id=3143 lang=golang
 *
 * [3143] Maximum Points Inside the Square
 */

// @lc code=start
func maxPointsInsideSquare(points [][]int, s string) int {
	res := 0
	check := func(r int) bool {
		apr := uint(0)
		for i, c := range s {
			if abs(points[i][0]) <= r && abs(points[i][1]) <= r {
				if apr&(1<<(byte(c)&31)) != 0 {
					return false
				}
				apr |= 1 << (byte(c) & 31)
			}
		}
		res = bits.OnesCount(apr)
		return true
	}
	lo, hi := 0, 1000_000_001
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return res
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// @lc code=end
