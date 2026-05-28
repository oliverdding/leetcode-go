package id3694

/*
 * @lc app=leetcode.cn id=3694 lang=golang
 *
 * [3694] Distinct Points Reachable After Substring Removal
 */

// @lc code=start
func distinctPoints(s string, k int) int {
	arr := []byte(s)
	x, y := 0, 0
	apprs := map[struct {
		x int
		y int
	}]struct{}{}

	for right, in := range arr {
		switch in {
		case 'U':
			y++
		case 'D':
			y--
		case 'L':
			x--
		case 'R':
			x++
		}

		left := right - k + 1
		if left < 0 {
			continue
		}

		apprs[struct {
			x int
			y int
		}{x, y}] = struct{}{}

		switch arr[left] {
		case 'U':
			y--
		case 'D':
			y++
		case 'L':
			x++
		case 'R':
			x--
		}
	}
	return len(apprs)
}

// @lc code=end
