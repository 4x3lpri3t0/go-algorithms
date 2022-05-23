package leetcode

import "testing"

func Test_countOdds(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: struct {
				low  int
				high int
			}{low: 1, high: 7}, want: 4,
		},
		{
			args: struct {
				low  int
				high int
			}{low: 3, high: 7}, want: 3,
		},
		{
			args: struct {
				low  int
				high int
			}{low: 3, high: 8}, want: 3,
		},
		{
			args: struct {
				low  int
				high int
			}{low: 8, high: 10}, want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOdds(tt.args.low, tt.args.high); got != tt.want {
				t.Errorf("countOdds() = %v, want %v", got, tt.want)
			}
		})
	}
}
