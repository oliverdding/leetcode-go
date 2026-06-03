package id2200

import "testing"

func Test_findKDistantIndices(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		key  int
		k    int
		want []int
	}{
		{
			"example 1",
			[]int{3, 4, 9, 1, 3, 9, 5},
			9,
			1,
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"example 2",
			[]int{2, 2, 2, 2, 2},
			2,
			2,
			[]int{0, 1, 2, 3, 4},
		},
		{
			"test 1",
			[]int{1, 1000, 1, 1000},
			1,
			1,
			[]int{0, 1, 2, 3},
		},
		{
			"test 2",
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			2,
			5,
			[]int{0, 1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findKDistantIndices(tt.nums, tt.key, tt.k)
			if !compareIntArray(got, tt.want) {
				t.Errorf("findKDistantIndices() = %v, want %v", got, tt.want)
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
