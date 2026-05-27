package id3439

import "testing"

func Test_maxFreeTime(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		eventTime int
		k         int
		startTime []int
		endTime   []int
		want      int
	}{
		{
			"example 1",
			5,
			1,
			[]int{1, 3},
			[]int{2, 5},
			2,
		},
		{
			"example 2",
			10,
			1,
			[]int{0, 2, 9},
			[]int{1, 4, 10},
			6,
		},
		{
			"example 3",
			5,
			2,
			[]int{0, 1, 2, 3, 4},
			[]int{1, 2, 3, 4, 5},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxFreeTime(tt.eventTime, tt.k, tt.startTime, tt.endTime)
			if got != tt.want {
				t.Errorf("maxFreeTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
