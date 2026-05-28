package id1652

/*
 * @lc app=leetcode.cn id=1652 lang=golang
 *
 * [1652] Defuse the Bomb
 */

// @lc code=start
func decrypt(code []int, k int) []int {
	s := len(code)
	right := k + 1
	if k < 0 {
		right = s
		k = -k
	}
	sum := 0
	for _, c := range code[right-k : right] {
		sum += c
	}
	res := make([]int, s)
	for i := 0; i < s; i++ {
		res[i] = sum
		sum += code[(right)%s] - code[(right-k)%s]
		right++
	}
	return res
}

// @lc code=end
