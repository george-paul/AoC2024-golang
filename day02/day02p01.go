package day02

import (
	"AoC_2024/util"
)

func checkIsSafe(diffs []int) (bool, int) {
	nPosDiffs := 0
	nNegDiffs := 0
	for _, num := range diffs {
		if num < 0 {
			nNegDiffs += 1
		}
		if num > 0 {
			nPosDiffs += 1
		}
	}

	isDec := nNegDiffs > nPosDiffs
	for idx := 0; idx < len(diffs); idx++ {
		diff := diffs[idx]
		if diff == 0 {
			return false, idx
		}
		if diff > 0 && isDec {
			return false, idx
		} else if diff < 0 && !isDec {
			return false, idx
		}
		if util.IntAbs(diff) < 1 || util.IntAbs(diff) > 3 {
			return false, idx
		}
	}
	return true, -1
}

func calcDiffs(line []int) []int {
	diffs := []int{}
	for nIdx := 0; nIdx < len(line)-1; nIdx++ {
		diffs = append(diffs, line[nIdx+1]-line[nIdx])
	}
	return diffs
}

func d02p01(reports [][]int) int {
	total := 0
	for _, line := range reports {
		diffs := calcDiffs(line)
		isSafe, _ := checkIsSafe(diffs)
		if isSafe {
			total += 1
		}
	}
	return total
}
