package id1888

/*
 * @lc app=leetcode.cn id=1888 lang=golang
 *
 * [1888] Minimum Number of Flips to Make the Binary String Alternating
 */

// @lc code=start
func minFlips(s string) int {
	k := len(s)
	cnt := uint(0) // t[i] != i % 2
	minOp := k
	for right := range 2*k - 1 {
		cnt += (uint(s[right%k]-'0') ^ uint(right)) & 1
		left := right - k + 1
		if left < 0 {
			continue
		}
		minOp = min(minOp, min(int(cnt), k-int(cnt)))
		cnt -= (uint(s[left]-'0') ^ uint(left)) & 1
	}
	return minOp
}

// @lc code=end
