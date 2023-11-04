package id278

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] First Bad Version
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	low, high := 0, n-1
	for low <= high {
		mid := low + (high-low)/2
		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}

// @lc code=end
