package id875

import "testing"

func Test_minEatingSpeed(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		piles []int
		h     int
		want  int
	}{
		{
			"example 1",
			[]int{3, 6, 7, 11},
			8,
			4,
		},
		{
			"example 2",
			[]int{30, 11, 23, 4, 20},
			5,
			30,
		},
		{
			"example 3",
			[]int{30, 11, 23, 4, 20},
			6,
			23,
		},
		{
			"test 1",
			[]int{1},
			5,
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minEatingSpeed(tt.piles, tt.h)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("minEatingSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}
