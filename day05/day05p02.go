package day05

func d05p02(orderings Orderings, updates Updates) int {
	total := 0
	for _, update := range updates {
		isInOrder, _, _ := checkUpdateOrder(orderings, update)
		if isInOrder {
			total += getMiddleNumber(update)
		}
	}

	return total
}
