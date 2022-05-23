package leetcode

func areAlmostEqual(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	diffCount := 0
	firstDiffIdx := -1
	secondDiffIdx := -1

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diffCount++

			if firstDiffIdx == -1 {
				firstDiffIdx = i
			} else {
				secondDiffIdx = i
			}
		}
	}

	return diffCount == 0 || (diffCount == 2 &&
		((s1[firstDiffIdx] == s2[secondDiffIdx]) &&
			(s1[secondDiffIdx] == s2[firstDiffIdx])))
}
