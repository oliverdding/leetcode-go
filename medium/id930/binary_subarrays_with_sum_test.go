package id930

import "testing"

func Test_solve(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		goal int
		want int
	}{
		{
			"example 1",
			[]int{0, 0, 0, 0, 0},
			0,
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solve(tt.nums, tt.goal)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("solve() = %v, want %v", got, tt.want)
			}
		})
	}
}
