package id374

/*
 * @lc app=leetcode.cn id=374 lang=golang
 *
 * [374] Guess Number Higher or Lower
 */

// @lc code=start
/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
	low, high := 0, n
	for low <= high {
		mid := low + (high-low)/2
		if tmp := guess(mid); tmp == -1 {
			high = mid - 1
		} else if tmp == 1 {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

// @lc code=end
