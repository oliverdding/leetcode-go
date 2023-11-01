package id33

import (
	"testing"
)

func Test_binarySearch(t *testing.T) {
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
			"single element (exists)",
			args{
				nums:   []int{1},
				target: 1,
			},
			0,
		},
		{
			"single element (non-exists)",
			args{
				nums:   []int{1},
				target: 2,
			},
			-1,
		},
		{
			name: "two elements (non-exists)",
			args: args{
				nums:   []int{1, 2},
				target: 2,
			},
			want: 1,
		},
		{
			"single element (non-exists)",
			args{
				nums:   []int{1, 2},
				target: 0,
			},
			-1,
		},
		{
			"simple",
			args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7},
				target: 5,
			},
			4,
		},
		{
			"the last element",
			args{
				nums:   []int{1, 2, 3, 4, 5, 6, 7},
				target: 7,
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("BinarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}

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
