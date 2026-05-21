package id1052

/*
 * @lc app=leetcode.cn id=1052 lang=golang
 *
 * [1052] Grumpy Bookstore Owner
 */

// @lc code=start
func maxSatisfied(customers []int, grumpy []int, minutes int) int {
	cnt := len(customers)
	normalSum := 0
	for i := range cnt {
		if grumpy[i] == 0 {
			normalSum += customers[i]
		}
	}
	grumpySum := 0
	grumpyMaxSum := 0
	for right := range cnt {
		if grumpy[right] == 1 {
			grumpySum += customers[right]
		}
		left := right - minutes + 1
		if left < 0 {
			continue
		}

		grumpyMaxSum = max(grumpyMaxSum, grumpySum)

		if grumpy[left] == 1 {
			grumpySum -= customers[left]
		}
	}
	return normalSum + grumpyMaxSum
}

// @lc code=end
