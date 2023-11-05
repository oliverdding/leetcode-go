package id892

import "testing"

func Test_getContactAreaCount(t *testing.T) {
	type args struct {
		grid [][]int
		i    int
		j    int
		high int
	}
	tests := []struct {
		name       string
		args       args
		wantResult int
	}{
		{
			"test1 0,0,1",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				0,
				0,
				1,
			},
			2,
		},
		{
			"test1 0,1,1",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				0,
				1,
				1,
			},
			3,
		},
		{
			"test1 0,1,2",
			args{
				[][]int{
					{1, 2},
					{3, 4},
				},
				0,
				1,
				2,
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotResult := getContactAreaCount(tt.args.grid, tt.args.i, tt.args.j, tt.args.high); gotResult != tt.wantResult {
				t.Errorf("getContactAreaCount() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
