package id3824

import "testing"

func Test_minimumK(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		nums []int
		want int
	}{
		// {
		// 	"example 1",
		// 	[]int{3, 7, 5},
		// 	3,
		// },
		// {
		// 	"example 2",
		// 	[]int{1},
		// 	1,
		// },
		{
			"test 1",
			[]int{1, 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumK(tt.nums)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minimumK() = %v, want %v", got, tt.want)
			}
		})
	}
}
