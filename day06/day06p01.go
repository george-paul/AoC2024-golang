package day06

// func isInBounds(r int, c int) bool {
// 	rInBounds := (r >= 0) && (r < dimH)
// 	cInBounds := (c >= 0) && (c < dimW)
// 	return rInBounds && cInBounds
// }

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
			curDir = (curDir + 1) % 4
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
