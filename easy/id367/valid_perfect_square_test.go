package id367

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				1,
			},
			want: true,
		},
		{
			"15",
			args{
				15,
			},
			false,
		},
		{
			"16",
			args{
				16,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.args.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
