package id69

import "testing"

func Test_mySqrt(t *testing.T) {
	type args struct {
		x int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{
				0,
			},
			0,
		},
		{
			"1",
			args{
				1,
			},
			1,
		},
		{
			"2",
			args{
				2,
			},
			1,
		},
		{
			"3",
			args{
				3,
			},
			1,
		},
		{
			"4",
			args{
				4,
			},
			2,
		}, {
			"8",
			args{
				8,
			},
			2,
		},
		{
			"9",
			args{
				9,
			},
			3,
		},
		{
			"25",
			args{
				25,
			},
			5,
		},
		{
			"26",
			args{
				26,
			},
			5,
		},
		{
			"80",
			args{
				80,
			},
			8,
		},
		{
			"82",
			args{
				82,
			},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mySqrt(tt.args.x); got != tt.want {
				t.Errorf("mySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
