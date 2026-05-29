package id3652

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		prices   []int
		strategy []int
		k        int
		want     int64
	}{
		{
			"test 3",
			[]int{38, 19, 19, 30, 22, 3},
			[]int{-1, -1, -1, 0, -1, -1},
			2,
			-25,
		},
		{
			"test 1",
			[]int{7, 1, 11, 22, 29},
			[]int{0, -1, -1, -1, -1},
			4,
			51,
		},
		{
			"test 2",
			[]int{12, 6, 14},
			[]int{-1, -1, 0},
			2,
			6,
		},
		{
			"example 1",
			[]int{4, 2, 8},
			[]int{-1, 0, 1},
			2,
			10,
		},
		{
			"example 2",
			[]int{5, 4, 3},
			[]int{1, 1, 0},
			2,
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices, tt.strategy, tt.k)
			if got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
