package id1011

import "testing"

func Test_shipWithinDays(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		weights []int
		days    int
		want    int
	}{
		{
			"test 61",
			[]int{1, 2, 3, 1, 1},
			4,
			3,
		},
		{
			"test 88",
			[]int{5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			8,
			10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shipWithinDays(tt.weights, tt.days)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("shipWithinDays() = %v, want %v", got, tt.want)
			}
		})
	}
}
