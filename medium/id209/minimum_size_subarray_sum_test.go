package id209

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		target int
		nums   []int
		want   int
	}{
		{
			"example 1",
			7,
			[]int{2,3,1,2,4,3},
			2,
		},
		{
			"example 2",
			4,
			[]int{1,4,4},
			1,
		},
		{
			"example 3",
			11,
			[]int{1,1,1,1,1,1,1,1},
			0,
		},
		{
			"failed 1",
			11,
			[]int{1,2,3,4,5},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
