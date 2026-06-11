package id2271

import "testing"

func Test_maximumWhiteTiles(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		tiles     [][]int
		carpetLen int
		want      int
	}{
		{
			"example 1",
			[][]int{{1, 5}, {10, 11}, {12, 18}, {20, 25}, {30, 32}},
			10,
			9,
		},
		{
			"example 2",
			[][]int{{10, 11}, {1, 1}},
			2,
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumWhiteTiles(tt.tiles, tt.carpetLen)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("maximumWhiteTiles() = %v, want %v", got, tt.want)
			}
		})
	}
}
