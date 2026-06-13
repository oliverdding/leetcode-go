package id2904

import "testing"

func Test_shortestBeautifulSubstring(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		k    int
		want string
	}{
		{
			"example 1",
			"100011001",
			3,
			"11001",
		},
		{
			"example 2",
			"1011",
			2,
			"11",
		},
		{
			"example 3",
			"000",
			1,
			"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shortestBeautifulSubstring(tt.s, tt.k)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("shortestBeautifulSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}
