package day05

func getMiddleNumber(update []int) int {
	midIdx := len(update) / 2
	return update[midIdx]
}

// returns true if update is in order
// returns false, a, b if update is not in order and orderings[a][b] was violated
func checkUpdateOrder(orderings Orderings, update []int) (bool, int, int) {
	beforeNums := []int{}
	for _, num := range update {
		for _, beforeNum := range beforeNums {
			if !orderings[beforeNum][num] {
				return false, beforeNum, num
			}
		}
		beforeNums = append(beforeNums, num)
	}
	return true, -1, -1
}

func d05p01(orderings Orderings, updates Updates) int {
	total := 0
	for _, update := range updates {
		isInOrder, _, _ := checkUpdateOrder(orderings, update)
		if isInOrder {
			total += getMiddleNumber(update)
		}
	}

	return total
}
