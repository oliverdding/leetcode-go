package id3439

/*
 * @lc app=leetcode.cn id=3439 lang=golang
 *
 * [3439] Reschedule Meetings for Maximum Free Time I
 */

// @lc code=start
func maxFreeTime(eventTime int, k int, startTime []int, endTime []int) int {
	m := len(startTime) + 1
	duringTime := make([]int, m)
	for i := 0; i < m; i++ {
		if i == 0 {
			duringTime[i] = startTime[i]
		} else if i == len(startTime) {
			duringTime[i] = eventTime - endTime[i-1]
		} else {
			duringTime[i] = startTime[i] - endTime[i-1]
		}
	}

	n := k + 1
	sum := 0
	res := 0
	for right, in := range duringTime {
		sum += in
		left := right - n + 1
		if left < 0 {
			continue
		}
		res = max(res, sum)
		sum -= duringTime[left]
	}
	return res
}

// @lc code=end
