package id3740

import "testing"

func Test_minimumDistance(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			"example 1",
			[]int{1, 2, 1, 1, 3},
			6,
		},
		{
			"example 2",
			[]int{1, 1, 2, 3, 2, 1, 2},
			8,
		},
		{
			"example 3",
			[]int{1},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumDistance(tt.nums)
			if got != tt.want {
				t.Errorf("minimumDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
