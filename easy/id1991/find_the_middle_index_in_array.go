package id1991

/*
 * @lc app=leetcode.cn id=1991 lang=golang
 *
 * [1991] Find the Middle Index in Array
 */

// @lc code=start
func findMiddleIndex(nums []int) int {
	sum := sumArray(nums)
	leftSum := 0
	for idx, num := range nums {
		if 2*leftSum+num == sum {
			return idx
		}
		leftSum += num
	}
	return -1
}

func sumArray(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// @lc code=end
