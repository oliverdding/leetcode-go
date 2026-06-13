package id1234

/*
 * @lc app=leetcode.cn id=1234 lang=golang
 *
 * [1234] Replace the Substring for Balanced String
 */

// @lc code=start
func balancedString(s string) int {
	n := len(s)
	m := n / 4
	cnt := make(map[byte]int, 4)
	for _, c := range s {
		cnt[byte(c)]++
	}
	if cnt['Q'] == m && cnt['W'] == m && cnt['E'] == m && cnt['R'] == m {
		return 0
	}
	left := 0
	res := n
	for right, in := range s {
		cnt[byte(in)]--
		for ; cnt['Q'] <= m && cnt['W'] <= m && cnt['E'] <= m && cnt['R'] <= m; left++ {
			res = min(res, right-left+1)
			cnt[s[left]]++
		}
	}
	return res
}

// @lc code=end
