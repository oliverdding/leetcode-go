package id1456

/*
 * @lc app=leetcode.cn id=1456 lang=golang
 *
 * [1456] Maximum Number of Vowels in a Substring of Given Length
 */

// @lc code=start
func maxVowels(s string, k int) int {
	nextCnt := 0
	res := 0

	// `i` stands for the right cornor of sliding window
	for right, in := range []byte(s) {
		if in == 'a' ||
			in == 'e' ||
			in == 'i' ||
			in == 'o' ||
			in == 'u' {
			nextCnt++
		}

		left := right - k + 1
		if left < 0 {
			// window not fill enough
			continue
		}

		// window is full
		res = max(res, nextCnt)
		if res == k {
			break
		}

		// remove the left cornor element for the next round
		out := s[left]
		if out == 'a' ||
			out == 'e' ||
			out == 'i' ||
			out == 'o' ||
			out == 'u' {
			nextCnt--
		}
	}

	return res
}

// @lc code=end
