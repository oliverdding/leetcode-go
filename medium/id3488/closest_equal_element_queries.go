package id3488

import "sort"

/*
 * @lc app=leetcode.cn id=3488 lang=golang
 *
 * [3488] Closest Equal Element Queries
 */

// @lc code=start
func solveQueries(nums []int, queries []int) []int {
	n := len(nums)
	idxs := make(map[int][]int)
	for i, num := range nums {
		if arr, ok := idxs[num]; !ok {
			idxs[num] = append(arr, 0, i)
		} else {
			idxs[num] = append(arr, i)
		}
	}
	for k, v := range idxs {
		v[0] = v[len(v)-1] - n
		v = append(v, v[1]+n)
		idxs[k] = v
	}
	res := make([]int, 0, len(queries))
	for _, query := range queries {
		arr := idxs[nums[query]]
		if len(arr) == 3 { // only one element
			res = append(res, -1)
			continue
		}
		idxOfIdx := sort.SearchInts(arr, query)
		res = append(res, min(arr[idxOfIdx+1]-query, query-arr[idxOfIdx-1]))
	}
	return res
}

// @lc code=end
