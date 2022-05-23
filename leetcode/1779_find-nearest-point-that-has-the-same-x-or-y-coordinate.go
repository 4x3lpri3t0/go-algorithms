package leetcode

import "math"

func nearestValidPoint(x int, y int, M [][]int) int {
	answer := -1
	minDistance := math.MaxInt
	for i := 0; i < len(M); i++ {
		if M[i][0] != x && M[i][1] != y {
			continue
		}

		// Manhattan Distance: |x1 - x2| + |y1 - y2|
		distance := int(math.Abs(float64(x-M[i][0])) + math.Abs(float64(y-M[i][1])))
		if distance < minDistance {
			minDistance = distance
			answer = i
		}
	}

	return answer
}

// Interesting alternative: https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/discuss/1096943/Golang-solution-100-100-with-explanation-and-images
