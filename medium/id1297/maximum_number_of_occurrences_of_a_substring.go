package id1297

/*
 * @lc app=leetcode.cn id=1297 lang=golang
 *
 * [1297] Maximum Number of Occurrences of a Substring
 */

// @lc code=start
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	res := 0
	strCnt := map[string]int{}
	charCnt := [26]int{}
	kinds := 0
	for right, in := range []byte(s) {
		if charCnt[in-'a'] == 0 {
			kinds++
		}
		charCnt[in-'a']++

		left := right - minSize + 1
		if left < 0 {
			continue
		}

		if kinds <= maxLetters {
			strCnt[s[left:right+1]]++
		}

		charCnt[s[left]-'a']--
		if charCnt[s[left]-'a'] == 0 {
			kinds--
		}
	}

	for _, v := range strCnt {
		res = max(res, v)
	}
	return res
}

// @lc code=end
