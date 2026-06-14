package id2062

import "testing"

func Test_countVowelSubstrings(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word string
		want int
	}{
		{
			"example 1",
			"aeiouu",
			2,
		},
		{
			"example 2",
			"unicornarihan",
			0,
		},
		{
			"example 3",
			"cuaieuouac",
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countVowelSubstrings(tt.word)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("countVowelSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
