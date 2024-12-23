package main

import "fmt"

func main() {
	day01()
}

func resultString(dayNum int, resultParts ...any) {
	fmt.Printf("---Day %02d---\n", dayNum)
	for idx, part := range resultParts {
		fmt.Print("Part ", idx+1, ":   ", part, "\n")
	}
}
