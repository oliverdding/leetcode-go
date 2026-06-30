package id1648

import "slices"

/*
 * @lc app=leetcode.cn id=1648 lang=golang
 *
 * [1648] Sell Diminishing-Valued Colored Balls
 */

// @lc code=start
const M = 1000_000_007

func maxProfit(inventory []int, orders int) int {
	res := 0
	check := func(b int) bool {
		sold, value, cnt := 0, 0, 0
		for _, inv := range inventory {
			if inv < b {
				continue
			}
			sold += inv - b + 1
			cnt += inv - b
			value += ((b + 1 + inv) * (inv - b) / 2) % M // count [b+1, inv] here
		}
		if sold < orders {
			return false
		}
		value += (b * (orders - cnt)) % M
		res = value % M
		return true
	}
	lo, hi := 0, slices.Max(inventory)+1
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

// @lc code=end
