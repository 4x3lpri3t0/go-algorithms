package leetcode

import (
	"math/bits"
)

func hammingWeight(num uint32) int {
	return bits.OnesCount32(num)
}

func hammingWeight_2(num uint32) int {
	count := uint32(0)
	for num > 0 {
		count += num % 2
		num >>= 1
	}
	return int(count)
}

func hammingWeight_3(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
