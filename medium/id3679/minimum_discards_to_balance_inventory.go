package id3679

/*
 * @lc app=leetcode.cn id=3679 lang=golang
 *
 * [3679]  Minimum Discards to Balance Inventory
 */

// @lc code=start
func minArrivalsToDiscard(arrivals []int, w int, m int) int {
	apprs := map[int]int{}
	added := make([]bool, len(arrivals))
	res := 0
	for i := 0; i < w; i++ {
		if apprs[arrivals[i]]+1 > m {
			res++
			continue
		}
		apprs[arrivals[i]]++
		added[i] = true
	}
	for i := w; i < len(arrivals); i++ {
		if added[i-w] {
			apprs[arrivals[i-w]]--
		}
		if apprs[arrivals[i]]+1 > m {
			res++
			continue
		}
		apprs[arrivals[i]]++
		added[i] = true
	}
	return res
}

// @lc code=end
