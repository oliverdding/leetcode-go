package id2799

/*
 * @lc app=leetcode.cn id=2799 lang=golang
 *
 * [2799] Count Complete Subarrays in an Array
 */

// @lc code=start
func countCompleteSubarrays(nums []int) int {
	allCnt := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		allCnt[num] = struct{}{}
	}
	cardix := len(allCnt)
	cnt := make(map[int]int, cardix)
	res, left := 0, 0
	for _, in := range nums {
		cnt[in]++
		for ; len(cnt) == cardix; left++ { // quit when condition not met
			out := nums[left]
			cnt[out]--
			if cnt[out] == 0 {
				delete(cnt, out)
			}
		}
		res += left
	}
	return res
}

// @lc code=end
