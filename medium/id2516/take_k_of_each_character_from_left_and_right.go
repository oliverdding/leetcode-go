package id2516

/*
 * @lc app=leetcode.cn id=2516 lang=golang
 *
 * [2516] Take K of Each Character From Left and Right
 */

// @lc code=start
func takeCharacters(s string, k int) int {
	length := [4]int{}
	for _, v := range s {
		length[v&31]++
	}
	if length[1] < k || length[2] < k || length[3] < k {
		return -1
	}
	res := 0
	left := 0
	count := [4]int{}
	for right, in := range s {
		count[in&31]++
		for ; count[1]+k > length[1] || count[2]+k > length[2] || count[3]+k > length[3]; left++ {
			count[s[left]&31]--
		}
		res = max(res, right-left+1)
	}
	return len(s) - res
}

// @lc code=end
