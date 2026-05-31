package id30

/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] Substring with Concatenation of All Words
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	apprs := make(map[string]int, len(words))
	for _, word := range words {
		apprs[word]++
	}

	res := []int{}
	for offset := 0; offset < len(words[0]); offset++ {
		tmp := findSubstringWithOffset(s, apprs, len(words), len(words[0]), offset)
		res = append(res, tmp...)
	}
	return res
}

func findSubstringWithOffset(s string, apprs map[string]int, wordsSize int, wordLen int, offset int) []int {
	res := []int{}
	localApprs := make(map[string]int)
	overload := 0
	for i := offset; i+wordLen <= len([]byte(s)); i += wordLen {
		if localApprs[s[i:i+wordLen]] == apprs[s[i:i+wordLen]] {
			overload++
		}
		localApprs[s[i:i+wordLen]]++
		left := i - (wordsSize-1)*wordLen
		if left < 0 {
			continue
		}
		if overload == 0 {
			res = append(res, left)
		}
		localApprs[s[left:left+wordLen]]--
		if localApprs[s[left:left+wordLen]] == apprs[s[left:left+wordLen]] {
			overload--
		}
	}
	return res
}

// @lc code=end
