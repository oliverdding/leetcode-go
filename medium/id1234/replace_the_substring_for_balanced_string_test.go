package id1234

import "testing"

func Test_balancedString(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		s    string
		want int
	}{
		// {
		// 	"example 1",
		// 	"QWER",
		// 	0,
		// },
		{
			"example 2",
			"QQWE",
			1,
		},
		{
			"example 3",
			"QQQW",
			2,
		},
		{
			"example 4",
			"QQQQ",
			3,
		},
		{
			"test 1",
			"WQWRQQQW",
			3,
		},
		{
			"test 2",
			"QEQRWRRWWWRQQQWQQEQEQREWRQEQRQQRRQEW",
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := balancedString(tt.s)
			// TODO: update the condition below to compare got with tt.want.
			if got != tt.want {
				t.Errorf("balancedString() = %v, want %v", got, tt.want)
			}
		})
	}
}
