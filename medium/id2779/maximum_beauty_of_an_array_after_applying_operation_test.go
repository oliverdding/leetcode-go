package id2779

import "testing"

func Test_maximumBeauty(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		want int
	}{
		{
			"example 1",
			[]int{4, 6, 1, 2},
			2,
			3,
		},
		{
			"example 2",
			[]int{1, 1, 1, 1},
			10,
			4,
		},
		{
			"test 1",
			[]int{5, 57, 46},
			15,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumBeauty(tt.nums, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("maximumBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}
