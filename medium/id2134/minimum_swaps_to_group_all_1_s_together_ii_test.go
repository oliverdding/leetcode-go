package id2134

import "testing"

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		{
			"example 1",
			[]int{0, 1, 0, 1, 1, 0, 0},
			1,
		},
		{
			"example 2",
			[]int{0, 1, 1, 1, 0, 0, 1, 1, 0},
			2,
		},
		{
			"example 3",
			[]int{1, 1, 0, 0, 1},
			0,
		},
		{
			"test 1",
			[]int{1, 1, 0, 1, 1},
			0,
		},
		{
			"test 2",
			[]int{1, 1, 1, 0, 0, 1, 0, 1, 1, 0},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minSwaps(tt.nums)
			if got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
