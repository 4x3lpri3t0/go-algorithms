package leetcode

import "testing"

func Test_subtractProductAndSum(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{args: struct{ n int }{n: 1}, want: 0},
		{args: struct{ n int }{n: 234}, want: 15},
		{args: struct{ n int }{n: 4421}, want: 21},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtractProductAndSum(tt.args.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
