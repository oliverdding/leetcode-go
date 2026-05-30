package id567

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] Permutation in String
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	mark := uint(0) // mark for letter
	apprs := make([]int, 26)
	k := len(s1)
	for _, char := range []byte(s1) {
		apprs[char-'a']++
		mark |= (1 << (char - 'a'))
	}
	for right, in := range []byte(s2) {
		apprs[in-'a']--
		if apprs[in-'a'] == 0 {
			mark &= ^(1 << (in - 'a'))
		} else {
			mark |= (1 << (in - 'a'))
		}
		left := right - k + 1
		if left < 0 {
			continue
		}
		if mark == 0 {
			return true
		}
		apprs[s2[left]-'a']++
		if apprs[s2[left]-'a'] == 0 {
			mark &= ^(1 << (s2[left] - 'a'))
		} else {
			mark |= (1 << (s2[left] - 'a'))
		}
	}
	return false
}

// @lc code=end
