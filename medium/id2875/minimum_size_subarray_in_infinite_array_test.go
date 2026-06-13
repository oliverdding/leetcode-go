package id2875

import "testing"

func Test_minSizeSubarray(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums   []int
		target int
		want   int
	}{
		{
			"test 1",
			[]int{2, 1, 5, 7, 7, 1, 6, 3},
			39,
			9,
		},
		{
			"example 1",
			[]int{1, 2, 3},
			5,
			2,
		},
		{
			"example 2",
			[]int{1, 1, 1, 2, 3},
			4,
			2,
		},
		{
			"example 3",
			[]int{2, 4, 6, 8},
			3,
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSizeSubarray(tt.nums, tt.target)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minSizeSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
