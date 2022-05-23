package leetcode

import "sort"

func Average_not_optimal(salary []int) float64 {
	sum := 0

	sort.Sort(sort.IntSlice(salary))

	for i := 1; i < len(salary)-1; i++ {
		sum += salary[i]
	}

	return float64(sum) / float64(len(salary)-2)
}

func average(salary []int) float64 {
	min, max, total := salary[0], salary[0], 0

	for _, v := range salary {
		total += v
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	return float64(total-min-max) / float64(len(salary)-2)
}
