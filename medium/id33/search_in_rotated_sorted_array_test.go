package id33

import (
	"testing"
)

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			4,
		},
		{
			"test2",
			args{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			-1,
		},
		{
			"test3",
			args{
				nums:   []int{1},
				target: 0,
			},
			-1,
		},
		{
			name: "unpassed1",
			args: args{
				nums:   []int{1},
				target: 1,
			},
			want: 0,
		},
		{
			name: "unpassed1",
			args: args{
				nums:   []int{5, 1, 3},
				target: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
