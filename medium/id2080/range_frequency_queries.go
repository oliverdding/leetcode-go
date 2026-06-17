package id2080

import "sort"

/*
 * @lc app=leetcode.cn id=2080 lang=golang
 *
 * [2080] Range Frequency Queries
 */

// @lc code=start
type RangeFreqQuery struct {
	data map[int][]int
}

func Constructor(arr []int) RangeFreqQuery {
	data := make(map[int][]int)
	for i, e := range arr {
		if list, ok := data[e]; ok {
			list = append(list, i)
			data[e] = list
		} else {
			data[e] = []int{i}
		}
	}
	return RangeFreqQuery{
		data,
	}
}

func (this *RangeFreqQuery) Query(left int, right int, value int) int {
	list := this.data[value]
	return sort.SearchInts(list, right+1) - sort.SearchInts(list, left)
}

/**
 * Your RangeFreqQuery object will be instantiated and called as such:
 * obj := Constructor(arr);
 * param_1 := obj.Query(left,right,value);
 */
// @lc code=end
