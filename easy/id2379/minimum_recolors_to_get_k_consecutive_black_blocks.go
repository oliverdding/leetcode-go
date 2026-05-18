package id2379

/*
 * @lc app=leetcode.cn id=2379 lang=golang
 *
 * [2379] Minimum Recolors to Get K Consecutive Black Blocks
 */

// @lc code=start
func minimumRecolors(blocks string, k int) int {
	whiteCnt := 0
	res := k
	for right, in := range []byte(blocks) {
		if in == 'W' {
			whiteCnt++
		}
		left := right - k + 1
		if left < 0 {
			continue
		}
		res = min(res, whiteCnt)
		if blocks[left] == 'W' {
			whiteCnt--
		}
	}
	return res
}

// @lc code=end
