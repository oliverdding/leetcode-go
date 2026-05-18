package id1423

/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] Maximum Points You Can Obtain from Cards
 */

// @lc code=start
func maxScore(cardPoints []int, k int) int {
	res := 0
	for i := range cardPoints[:k] {
		res += cardPoints[i]
	}

	sum := res
	for i := 1; i <= k; i++ {
		sum += cardPoints[len(cardPoints)-i] - cardPoints[k-i]
		res = max(res, sum)
	}
	return res
}

// @lc code=end
