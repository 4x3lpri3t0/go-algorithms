package leetcode

import "math"

func getSlope(p1, p2 []int) float64 {
	denominator := (p2[0] - p1[0])

	if denominator == 0 {
		return math.Inf(1)
	}

	divisor := (p1[1] - p2[1])

	return float64(divisor) / float64(denominator)
}

func checkStraightLine(coordinates [][]int) bool {
	if len(coordinates) == 2 {
		return true
	}

	slope := getSlope(coordinates[0], coordinates[1])

	for i := 2; i < len(coordinates); i++ {
		if getSlope(coordinates[i], coordinates[i-1]) != slope {
			return false
		}
	}

	return true
}
