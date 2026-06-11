package id2271

import "slices"

/*
 * @lc app=leetcode.cn id=2271 lang=golang
 *
 * [2271] Maximum White Tiles Covered by a Carpet
 */

// @lc code=start
func maximumWhiteTiles(tiles [][]int, carpetLen int) int {
	slices.SortFunc(tiles, func(a, b []int) int {
		return a[0] - b[0]
	})
	res := 0
	left := 0
	cover := 0
	for _, tile := range tiles {
		tl, tr := tile[0], tile[1]
		cover += tr - tl + 1

		carpetLeft := tr - carpetLen + 1
		for ; tiles[left][1] < carpetLeft; left++ {
			cover -= tiles[left][1] - tiles[left][0] + 1
		}

		uncover := max(0, carpetLeft-tiles[left][0])
		res = max(res, cover-uncover)
	}
	return res
}

// @lc code=end
