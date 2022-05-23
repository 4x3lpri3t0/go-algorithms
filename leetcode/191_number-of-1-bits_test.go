package leetcode

import "testing"

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: struct{ num uint32 }{num: 0}, want: 0,
		},
		{
			args: struct{ num uint32 }{num: 00000000000000000000000000001011}, want: 3,
		},
		{
			args: struct{ num uint32 }{num: 00000000000000000000000010000000}, want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
