package id438

/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] Find All Anagrams in a String
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	apprs := make([]int, 26)
	mark := uint(0)
	for _, char := range []byte(p) {
		apprs[char-'a']++
		if apprs[char-'a'] != 0 {
			mark |= (1 << (char - 'a'))
		}
	}
	k := len(p)
	res := []int{}
	for right, in := range []byte(s) {
		apprs[in-'a']--
		if apprs[in-'a'] == 0 {
			mark &= ^(1 << (in - 'a'))
		} else {
			mark |= (1 << (in - 'a'))
		}
		// if window is not full
		left := right - k + 1
		if left < 0 {
			continue
		}
		// judge result
		if mark == 0 {
			res = append(res, left)
		}
		out := []byte(s)[left]
		// prepare for next round
		apprs[out-'a']++
		if apprs[out-'a'] == 0 {
			mark &= ^(1 << (out - 'a'))
		} else {
			mark |= (1 << (out - 'a'))
		}
	}
	return res
}

// @lc code=end
