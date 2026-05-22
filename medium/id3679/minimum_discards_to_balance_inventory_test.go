package id3679

import "testing"

func Test_minArrivalsToDiscard(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		arrivals []int
		w        int
		m        int
		want     int
	}{
		{
			"test 1",
			[]int{7, 10, 6, 9, 2, 8, 7, 4, 3},
			5,
			1,
			0,
		},
		{
			"test 2",
			[]int{1, 2, 3, 3, 3, 4},
			3,
			2,
			1,
		},
		{
			"test 3",
			[]int{4, 8, 1, 4, 6, 3, 1, 1, 4, 4},
			9,
			2,
			2,
		},
		{
			"test 4",
			[]int{7, 3, 9, 9, 7, 3, 5, 9, 7, 2, 6, 10, 9, 7, 9, 1, 3, 6, 2, 4, 6, 2, 6, 8, 4, 8, 2, 7, 5, 6},
			10,
			1,
			13,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minArrivalsToDiscard(tt.arrivals, tt.w, tt.m)
			if got != tt.want {
				t.Errorf("minArrivalsToDiscard() = %v, want %v", got, tt.want)
			}
		})
	}
}
