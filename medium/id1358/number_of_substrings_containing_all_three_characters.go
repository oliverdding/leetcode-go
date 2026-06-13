package id1358

/*
 * @lc app=leetcode.cn id=1358 lang=golang
 *
 * [1358] Number of Substrings Containing All Three Characters
 */

// @lc code=start
func numberOfSubstrings(s string) int {
	cnt := [4]int{}
	res := 0
	left := 0
	for _, in := range s {
		cnt[in&31]++
		for ; cnt[1] > 0 && cnt[2] > 0 && cnt[3] > 0; left++ {
			cnt[s[left]&31]--
		}
		res += left // 0~left-1 are all okay
	}
	return res
}

// @lc code=end
