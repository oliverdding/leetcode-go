package id904

/*
 * @lc app=leetcode.cn id=904 lang=golang
 *
 * [904] Fruit Into Baskets
 */

// @lc code=start
func totalFruit(fruits []int) int {
	res := 0
	left := 0
	cnt := make(map[int]int, 3)
	for right, in := range fruits {
		cnt[in]++
		for ; len(cnt) == 3; left++ {
			cnt[fruits[left]]--
			if cnt[fruits[left]] == 0 {
				delete(cnt, fruits[left])
			}
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
