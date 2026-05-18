package id2379

import "testing"

func Test_minimumRecolors(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		blocks string
		k      int
		want   int
	}{
		{
			"test 1",
			"W",
			1,
			1,
		},
		{
			"test 2",
			"WW",
			2,
			2,
		},
		{
			"test 3",
			"BB",
			2,
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minimumRecolors(tt.blocks, tt.k)
			if got != tt.want {
				t.Errorf("minimumRecolors() = %v, want %v", got, tt.want)
			}
		})
	}
}
