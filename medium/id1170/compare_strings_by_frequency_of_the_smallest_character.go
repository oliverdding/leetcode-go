package id1170

import "slices"

/*
 * @lc app=leetcode.cn id=1170 lang=golang
 *
 * [1170] Compare Strings by Frequency of the Smallest Character
 */

// @lc code=start
func numSmallerByFrequency(queries []string, words []string) []int {
	arr2 := make([]int, 0, len(words))
	for _, word := range words {
		arr2 = append(arr2, f(word))
	}
	slices.Sort(arr2)
	res := make([]int, 0, len(queries))
	for _, query := range queries {
		cnt := f(query)
		res = append(res, len(words)-lowerBound(arr2, cnt+1))
	}
	return res
}

func lowerBound(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if nums[mid] >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

func f(s string) int {
	min := 'z'
	cnt := 0
	for _, c := range s {
		switch {
		case c < min:
			min = c
			cnt = 1
		case c == min:
			cnt++
		}
	}
	return cnt
}

// @lc code=end
