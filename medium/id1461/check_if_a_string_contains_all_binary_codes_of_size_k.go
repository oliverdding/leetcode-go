package id1461

/*
 * @lc app=leetcode.cn id=1461 lang=golang
 *
 * [1461] Check If a String Contains All Binary Codes of Size K
 */

// @lc code=start
func hasAllCodes(s string, k int) bool {
	arr := []byte(s)
	apprs := map[int]struct{}{}
	binaryNum := 0
	mask := ^(1 << (k - 1))
	cnt := 1 << k
	for i := range arr {
		binaryNum = (binaryNum&mask)<<1 | int(arr[i]&1)
		left := i - k + 1
		if left < 0 {
			continue
		}
		apprs[binaryNum] = struct{}{}
		if len(apprs) == cnt {
			return true
		}
		if cnt-len(apprs) > len(arr)-i-1 {
			break
		}
	}
	return false
}

// @lc code=end
