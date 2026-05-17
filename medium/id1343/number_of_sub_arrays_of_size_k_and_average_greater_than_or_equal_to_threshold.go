package id1343

/*
 * @lc app=leetcode.cn id=1343 lang=golang
 *
 * [1343] Number of Sub-arrays of Size K and Average Greater than or Equal to Threshold
 */

// @lc code=start
func numOfSubarrays(arr []int, k int, threshold int) int {
	sum := 0
	cnt := 0
	for right, in := range arr {
		sum += in
		left := right - k + 1
		if left < 0 {
			continue
		}
		if float64(sum)/float64(k) >= float64(threshold) {
			cnt++
		}
		sum -= arr[left]
	}
	return cnt
}

// @lc code=end
