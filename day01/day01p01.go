package day01

import (
	"AoC_2024/util"
	"sort"
)

func d01p01(l1 [1000]int, l2 [1000]int) int {
	sort.Ints(l1[:])
	sort.Ints(l2[:])

	total := 0

	for idx := range l1 {
		total += util.IntAbs(l1[idx] - l2[idx])
	}

	return total
}
