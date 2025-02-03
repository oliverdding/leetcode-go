package id48

/*
 * @lc app=leetcode.cn id=48 lang=golang
 *
 * [48] Rotate Image
 */

// Example 2: length is 4
// -> (0,2) = (1,0)
// -> (1,0) = (3,1)
// -> (3,1) = (2,3)
// -> (2,3) = (0,2)

// @lc code=start
func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-1-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = tmp
		}
	}
}

// @lc code=end
