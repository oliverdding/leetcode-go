package id2300

import "slices"

/*
 * @lc app=leetcode.cn id=2300 lang=golang
 *
 * [2300] Successful Pairs of Spells and Potions
 */

// @lc code=start
func successfulPairs(spells []int, potions []int, success int64) []int {
	slices.Sort(potions)
	res := make([]int, 0, len(spells))
	for _, spell := range spells {
		res = append(res, len(potions)-lowerBound(potions, spell, success))
	}
	return res
}

func lowerBound(nums []int, spell int, success int64) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)>>1
		if int64(nums[mid])*int64(spell) >= success {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}

// @lc code=end
