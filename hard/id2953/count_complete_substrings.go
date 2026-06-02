package id2953

/*
 * @lc app=leetcode.cn id=2953 lang=golang
 *
 * [2953] Count Complete Substrings
 */

// @lc code=start
func countCompleteSubstrings(word string, k int) int {
	res := 0
	subWordIndex := []int{0}
	for i := 1; i < len([]byte(word)); i++ {
		if int([]byte(word)[i])-int([]byte(word)[i-1]) > 2 || int([]byte(word)[i])-int([]byte(word)[i-1]) < -2 {
			subWordIndex = append(subWordIndex, i)
		}
	}
	subWordIndex = append(subWordIndex, len([]byte(word)))

	for n := 1; n <= 26 && n*k <= len([]byte(word)); n++ {
		for i := 0; i < len(subWordIndex)-1; i++ {
			if n*k <= subWordIndex[i+1]-subWordIndex[i] { // substring size is larger than real k
				apprs := [26]int{}
				check := func() {
					for _, v := range apprs {
						if v != 0 && v != k {
							return
						}
					}
					res++
				}
				subWord := []byte(word[subWordIndex[i]:subWordIndex[i+1]])
				for right, in := range subWord {
					apprs[in-'a']++
					if left := right - n*k + 1; left >= 0 {
						check()
						apprs[subWord[left]-'a']--
					}
				}
			}
		}
	}
	return res
}

// @lc code=end
