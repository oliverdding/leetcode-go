package id29

/*
 * @lc app=leetcode.cn id=29 lang=golang
 *
 * [29] Divide Two Integers
 */

// @lc code=start
import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if dividend == math.MinInt && divisor == -1 {
		return math.MaxInt
	}

}

func myPlus()

// @lc code=end
