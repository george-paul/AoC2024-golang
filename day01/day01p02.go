package day01

func d01p02(l1 [1000]int, l2 [1000]int) int {
	valMap := map[int]int{}

	for _, val := range l1 {
		valMap[val] = 0
	}

	for _, val := range l2 {
		v, ok := valMap[val]
		if ok {
			valMap[val] = v + 1
		}
	}

	total := 0
	for key, val := range valMap {
		total += key * val
	}

	return total
}
