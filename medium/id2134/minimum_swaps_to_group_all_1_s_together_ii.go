package id2134

/*
 * @lc app=leetcode.cn id=2134 lang=golang
 *
 * [2134] Minimum Swaps to Group All 1's Together II
 */

// @lc code=start
func minSwaps(nums []int) int {
	size := len(nums)
	k := 0
	for _, num := range nums {
		k += num
	}
	if k == 0 || k == size {
		return 0
	}

	cntOne := 0
	maxOne := 0
	for i := 0; i <= size+k-2; i++ {
		cntOne += nums[i%size]
		left := i - k + 1
		if left < 0 {
			continue
		}
		maxOne = max(maxOne, cntOne)
		cntOne -= nums[left]
	}

	return k - maxOne
}

// @lc code=end
