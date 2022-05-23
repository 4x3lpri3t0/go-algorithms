package leetcode

func countOdds(low int, high int) int {
	diff := (high - low) / 2

	if low%2 != 0 || high%2 != 0 {
		diff += 1
	}

	return diff
}
