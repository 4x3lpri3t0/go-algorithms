package leetcode

import "testing"

func Test_checkStraightLine(t *testing.T) {
	type args struct {
		coordinates [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{[][]int{{1, 2}, {6, 7}}}, true},
		{"2", args{[][]int{{1, 1}, {3, 4}}}, true},
		{"3", args{[][]int{{0, 0}, {0, 1}, {0, -1}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkStraightLine(tt.args.coordinates); got != tt.want {
				t.Errorf("checkStraightLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
