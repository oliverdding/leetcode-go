package id892

/*
 * @lc app=leetcode.cn id=892 lang=golang
 *
 * [892] Surface Area of 3D Shapes
 */

// @lc code=start
func surfaceArea(grid [][]int) int {
	n := len(grid)
	cubes := 0
	contactArea := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cubes += grid[i][j]
			for high := 1; high <= grid[i][j]; high++ {
				contactArea += getContactAreaCount(grid, i, j, high)
			}
		}
	}
	return cubes*6 - contactArea
}

func getContactAreaCount(grid [][]int, i, j, high int) (result int) {
	if i > 0 && grid[i-1][j] >= high {
		result++
	}
	if i < len(grid)-1 && grid[i+1][j] >= high {
		result++
	}
	if j > 0 && grid[i][j-1] >= high {
		result++
	}
	if j < len(grid)-1 && grid[i][j+1] >= high {
		result++
	}
	if high > 1 {
		result++
	}
	if high < grid[i][j] {
		result++
	}
	return result
}

// @lc code=end
