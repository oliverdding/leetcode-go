package id3795

import "testing"

func Test_minLength(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		{
			"example 1",
			[]int{2, 2, 3, 1},
			4,
			2,
		},
		{
			"example 2",
			[]int{3, 2, 3, 4},
			5,
			2,
		},
		{
			"example 3",
			[]int{5, 5, 4},
			5,
			1,
		},
		{
			"test 1",
			[]int{1, 12},
			7,
			1,
		},
		{
			"test 2",
			[]int{154, 163, 156, 47, 151, 151, 22, 203, 149, 130, 115, 71, 3, 14, 30, 132, 208, 100, 125, 160, 36, 103, 125, 51, 187, 137, 157, 217, 53, 4, 146, 14, 20, 59, 224, 129, 161, 29},
			1247,
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minLength(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
