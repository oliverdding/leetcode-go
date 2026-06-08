package id2024

/*
 * @lc app=leetcode.cn id=2024 lang=golang
 *
 * [2024] Maximize the Confusion of an Exam
 */

// @lc code=start
func maxConsecutiveAnswers(answerKey string, k int) int {
	res := 0
	left := 0
	cnt := [2]int{}
	for right, in := range answerKey {
		cnt[in>>1&1]++
		for ; cnt[0] > k && cnt[1] > k; left++ {
			cnt[answerKey[left]>>1&1]--
		}
		res = max(res, right-left+1)
	}
	return res
}

// @lc code=end
