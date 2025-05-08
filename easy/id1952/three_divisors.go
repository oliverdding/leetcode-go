package id1952

import "math"

/*
 * @lc app=leetcode.cn id=1952 lang=golang
 *
 * [1952] Three Divisors
 */

// @lc code=start
func isThree(n int) bool {
	if n == 1 {
		return false
	}
	// only n*n number can be
	tmp := math.Sqrt(float64(n))
	if tmp-float64(int(tmp)) != 0 {
		return false
	}
	// [2, \sqrt{n})
	for i := 2; i < int(tmp); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// @lc code=end
