package day02

import (
	"fmt"
	"slices"
)

func d02p02(reports [][]int) int {
	total := 0
	for lIdx, line := range reports {
		diffs := calcDiffs(line)
		isSafe, atIdx := checkIsSafe(diffs)
		if isSafe {
			total += 1
			fmt.Print(lIdx, "\n")
		} else {
			newLine1 := append([]int{}, line...)
			newLine1 = slices.Delete(newLine1, atIdx, atIdx+1)
			newDiffs1 := calcDiffs(newLine1)
			isNew1Safe, _ := checkIsSafe(newDiffs1)
			if isNew1Safe {
				total += 1
				fmt.Print(lIdx, "\n")
				continue
			}
			newLine2 := append([]int{}, line...)
			newLine2 = slices.Delete(newLine2, atIdx+1, atIdx+2)
			newDiffs2 := calcDiffs(newLine2)
			isNew2Safe, _ := checkIsSafe(newDiffs2)
			if isNew2Safe {
				total += 1
				fmt.Print(lIdx, "\n")
				continue
			}
		}
	}
	return total
}
