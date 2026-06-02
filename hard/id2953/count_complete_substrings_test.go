package id2953

import "testing"

func Test_countCompleteSubstrings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word string
		k    int
		want int
	}{
		{
			"example 1",
			"igigee",
			2,
			3,
		},
		{
			"example 2",
			"aaabbbccc",
			3,
			6,
		},
		{
			"test 1",
			"aaa",
			1,
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countCompleteSubstrings(tt.word, tt.k)
			if got != tt.want {
				t.Errorf("countCompleteSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
