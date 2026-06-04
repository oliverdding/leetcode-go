package id2653

import "testing"

func Test_getSubarrayBeauty(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		k    int
		x    int
		want []int
	}{
		{
			"example 1",
			[]int{1, -1, -3, -2, 3},
			3,
			2,
			[]int{-1, -2, -2},
		},
		{
			"test 1",
			[]int{-1, -2, -3, -4, -5},
			2,
			2,
			[]int{-1, -2, -3, -4},
		},
		{
			"test 2",
			[]int{-3, 1, 2, -3, 0, -3},
			2,
			1,
			[]int{-3, 0, -3, -3, -3},
		},
		{
			"test 3",
			[]int{5},
			1,
			1,
			[]int{0},
		},
		{
			"test 4",
			[]int{-4, 0, -1},
			2,
			2,
			[]int{0, 0},
		},
		{
			"test 5",
			[]int{26, -36, 11, -47, 32},
			3,
			2,
			[]int{0, -36, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSubarrayBeauty(tt.nums, tt.k, tt.x)
			// TODO: update the condition below to compare got with tt.want.
			if !compareIntArray(got, tt.want) {
				t.Errorf("getSubarrayBeauty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func compareIntArray(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
