package id2062

/*
 * @lc app=leetcode.cn id=2062 lang=golang
 *
 * [2062] Count Vowel Substrings of a String
 */

// @lc code=start
func countVowelSubstrings(word string) int {
	res, left, lastVowel := 0, 0, 0
	cnt := make(map[byte]int, 5)
	for right, in := range word {
		if !isVowel(byte(in)) {
			clear(cnt)
			lastVowel = right + 1
			left = right + 1
			continue
		}
		cnt[byte(in)]++
		for ; cnt['a'] >= 1 && cnt['e'] >= 1 && cnt['i'] >= 1 && cnt['o'] >= 1 && cnt['u'] >= 1; left++ { // quit when condition not met
			cnt[word[left]]--
		}
		res += left - lastVowel
	}
	return res
}

func isVowel(c byte) bool {
	if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		return true
	}
	return false
}

// @lc code=end
