package day06

func d06p01(grid Grid, begR int, begC int) int {
	r := begR
	c := begC
	// count starting cell
	total := 1
	grid[begR][begC] = 3

	curDir := uint8(cuu) // initial direction is up
	for {
		newR, newC, err := moveCellIndex(curDir, r, c)
		next := grid[newR][newC]
		if err != nil {
			break
		}
		if next == 1 {
			curDir = (curDir % 4) + 1
		} else {
			if next != 3 {
				grid[newR][newC] = 3
				total++
			}
			r = newR
			c = newC
		}
	}
	return total
}
