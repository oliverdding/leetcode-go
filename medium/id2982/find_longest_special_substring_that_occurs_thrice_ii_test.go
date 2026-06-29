package id2982

import "testing"

func Test_maximumLength(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		{
			"example 1",
			"aaaa",
			2,
		},
		{
			"example 2",
			"abcdef",
			-1,
		},
		{
			"example 3",
			"abcaba",
			1,
		},
		{
			"test 1",
			"akphhppppp",
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumLength(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("maximumLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
