package id1423

import "testing"

func Test_maxScore(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		cardPoints []int
		k          int
		want       int
	}{
		{
			"example 1",
			[]int{1, 2, 3, 4, 5, 6, 1},
			3,
			12,
		},
		{
			"example 3",
			[]int{9, 7, 7, 9, 7, 7, 9},
			7,
			55,
		},
		{
			"example 4",
			[]int{1, 1000, 1},
			1,
			1,
		},
		{
			"example 5",
			[]int{1, 79, 80, 1, 1, 1, 200, 1},
			3,
			202,
		},
		{
			"test 5",
			[]int{111, 111, 111, 1, 1, 1, 1, 1},
			3,
			333,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxScore(tt.cardPoints, tt.k)
			if got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
