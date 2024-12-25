package util

import "fmt"

func ResultString(dayNum int, resultParts ...any) {
	fmt.Printf("---Day %02d---\n", dayNum)
	for idx, part := range resultParts {
		fmt.Print("Part ", idx+1, ":   ", part, "\n")
	}
}

func IntAbs(num int) int {
	if num <= 0 {
		return -num
	} else {
		return num
	}
}
