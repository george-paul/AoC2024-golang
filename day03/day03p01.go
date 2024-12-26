package day03

import (
	"regexp"
	"strconv"
	"strings"
)

func doMul(mulString string) int {
	mulStringLen := len(mulString)
	numsString := mulString[4 : mulStringLen-1]
	nums := strings.Split(numsString, ",")
	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1 * n2
}

func d03p01(inputString string) int {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	allMulStrings := re.FindAllString(inputString, -1)
	total := 0
	for _, mulString := range allMulStrings {
		total += doMul(mulString)
	}

	return total
}
