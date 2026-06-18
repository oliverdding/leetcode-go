package id3488

import "testing"

func Test_solveQueries(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums    []int
		queries []int
		want    []int
	}{
		{
			"exmaple 1",
			[]int{1, 3, 1, 4, 1, 3, 2},
			[]int{0, 3, 5},
			[]int{2, -1, 3},
		},
		{
			"exmaple 2",
			[]int{1, 2, 3, 4},
			[]int{0, 1, 2, 3},
			[]int{-1, -1, -1, -1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solveQueries(tt.nums, tt.queries)
			// TODO: update the condition below to compare got with tt.want.
			if !compareIntArray(got, tt.want) {
				t.Errorf("solveQueries() = %v, want %v", got, tt.want)
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
