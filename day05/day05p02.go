package day05

func swapNumsInUpdate(update []int, a int, b int) []int {
	newUpdate := update
	temp := update[a]
	update[a] = update[b]
	update[b] = temp
	return newUpdate
}

func d05p02(orderings Orderings, updates Updates) int {
	total := 0
	for _, update := range updates {
		isInOrder, a, b := checkUpdateOrder(orderings, update)
		for !isInOrder {
			update = swapNumsInUpdate(update, a, b)
			isInOrder, a, b = checkUpdateOrder(orderings, update)
			if isInOrder {
				total += getMiddleNumber(update)
				break
			}
		}
	}
	return total
}
