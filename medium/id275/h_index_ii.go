package id275

/*
 * @lc app=leetcode.cn id=275 lang=golang
 *
 * [275] H-Index II
 */

// @lc code=start
func hIndex(citations []int) int {
	check := func(h int) bool {
		return citations[len(citations)-h] >= h
	}
	lo, hi := 1, len(citations)+1
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if check(mid) {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi - 1
}

// @lc code=end
