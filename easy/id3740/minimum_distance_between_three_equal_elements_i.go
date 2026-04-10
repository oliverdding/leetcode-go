package id3740

/*
 * @lc app=leetcode.cn id=3740 lang=golang
 *
 * [3740] Minimum Distance Between Three Equal Elements I
 */

// @lc code=start
func minimumDistance(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	// count appearance
	cntMap := make(map[int]int, 100)
	for _, num := range nums {
		cntMap[num]++
	}
	// filter out >= 3
	idxMap := make(map[int][]int)
	for key, val := range cntMap {
		if val < 3 {
			continue
		}
		idxMap[key] = make([]int, 0, val)
	}
	// record it's indexes
	for idx, num := range nums {
		if arr, ok := idxMap[num]; ok {
			arr = append(arr, idx)
			idxMap[num] = arr
		}
	}
	// calculate the minimous
	res := 101
	for _, val := range idxMap {
		for i := 0; i <= len(val)-3; i++ {
			res = min(res, val[i+2]-val[i])
		}
	}
	if res == 101 {
		return -1
	}
	return res * 2
}

// @lc code=end
