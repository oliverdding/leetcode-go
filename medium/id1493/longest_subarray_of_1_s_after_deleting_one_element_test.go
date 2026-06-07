package id1493

import "testing"

func Test_longestSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			"example 1",
			[]int{1, 1, 0, 1},
			3,
		},
		{
			"example 2",
			[]int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			5,
		},
		{
			"example 3",
			[]int{1, 1, 1},
			2,
		},
		{
			"test 1",
			[]int{1, 1, 1, 0, 1, 0, 1, 1, 1, 1, 0},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestSubarray(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
