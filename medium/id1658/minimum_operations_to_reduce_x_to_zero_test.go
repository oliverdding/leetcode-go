package id1658

import "testing"

func Test_minOperations(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		x    int
		want int
	}{
		{
			"example 1",
			[]int{1, 1, 4, 2, 3},
			5,
			2,
		},
		{
			"example 2",
			[]int{5, 6, 7, 8, 9},
			4,
			-1,
		},
		{
			"example 3",
			[]int{3, 2, 20, 1, 1, 3},
			10,
			5,
		},
		{
			"test 1",
			[]int{1, 1, 1, 1, 1},
			5,
			5,
		},
		{
			"test 2",
			[]int{1, 1, 4, 1, 1},
			6,
			3,
		},
		{
			"test 3",
			[]int{5, 1, 4, 2, 3},
			6,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.x)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
