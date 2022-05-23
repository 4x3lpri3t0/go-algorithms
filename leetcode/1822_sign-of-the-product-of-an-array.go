package leetcode

func arraySign(nums []int) int {
	var sign int
	for _, v := range nums {
		if v == 0 {
			return 0
		}
		if v < 0 {
			sign++
		}
	}

	if sign%2 == 0 {
		return 1
	}

	return -1
}
