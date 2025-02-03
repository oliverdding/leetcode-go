package id498

/*
 * @lc app=leetcode.cn id=498 lang=golang
 *
 * [498] Diagonal Traverse
 */

// @lc code=start
func findDiagonalOrder(mat [][]int) []int {
	m, n := len(mat), len(mat[0])
	ans := make([]int, 0, m*n)
	for i := 0; i < m+n-1; i++ {
		var (
			x int
			y int
		)
		if i%2 == 1 {
			if i < n {
				x = 0
				y = i
			} else {
				x = i - n + 1
				y = n - 1
			}
			for x < m && y >= 0 {
				ans = append(ans, mat[x][y])
				x++
				y--
			}
		} else {
			if i < m {
				x = i
				y = 0
			} else {
				x = m - 1
				y = i - m + 1
			}
			for x >= 0 && y < n {
				ans = append(ans, mat[x][y])
				x--
				y++
			}
		}
	}
	return ans
}

// @lc code=end
