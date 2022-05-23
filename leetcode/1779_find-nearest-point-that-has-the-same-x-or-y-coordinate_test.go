package leetcode

import "testing"

func Test_nearestValidPoint(t *testing.T) {
	type args struct {
		x      int
		y      int
		points [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{1, 1, [][]int{{0, 0}, {2, 2}, {3, 3}}}, -1},
		{"2", args{3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}}, 2},
		{"3", args{3, 4, [][]int{{3, 4}}}, 0},
		{"4", args{3, 4, [][]int{{2, 3}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nearestValidPoint(tt.args.x, tt.args.y, tt.args.points); got != tt.want {
				t.Errorf("nearestValidPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
