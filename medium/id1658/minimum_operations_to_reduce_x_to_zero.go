package id1658

/*
 * @lc app=leetcode.cn id=1658 lang=golang
 *
 * [1658] Minimum Operations to Reduce X to Zero
 */

// @lc code=start
func minOperations(nums []int, x int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum < x {
		return -1
	}
	size := -1
	left := 0
	s := 0
	for right, in := range nums {
		s += in
		for ; s > sum-x; left++ {
			s -= nums[left]
		}
		if s == sum-x {
			size = max(size, right-left+1)
		}
	}
	if size == -1 {
		return -1
	}
	return len(nums) - size
}

// @lc code=end
