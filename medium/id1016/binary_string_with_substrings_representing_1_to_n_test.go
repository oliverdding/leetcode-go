package id1016

import "testing"

func Test_queryString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		n    int
		want bool
	}{
		{
			"example 1",
			"0110",
			3,
			true,
		},
		{
			"example 2",
			"0110",
			4,
			false,
		},
		{
			"test 1",
			"1111000101",
			5,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := queryString(tt.s, tt.n)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("queryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
