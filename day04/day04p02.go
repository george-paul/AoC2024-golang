package day04

func checkA(rIdx int, cIdx int, cells Cells) (isXMAS bool) {
	var newRIdx int
	var newCIdx int

	corners := [4]uint8{0, 0, 0, 0}

	for idx, dir := range []uint8{cul, cur, cdl, cdr} {
		newRIdx, newCIdx, _ = moveCellIndex(dir, rIdx, cIdx)
		newCell := cells[newRIdx][newCIdx]
		corners[idx] = newCell.letter
	}

	valid0 := [4]uint8{2, 4, 2, 4}
	valid1 := [4]uint8{2, 2, 4, 4}
	valid2 := [4]uint8{4, 2, 4, 2}
	valid3 := [4]uint8{4, 4, 2, 2}

	switch corners {
	case valid0, valid1, valid2, valid3:
		return true
	}

	return false
}

func d04p02(cells Cells) int {
	total := 0
	for rIdx, line := range cells {
		if rIdx == 0 || rIdx == dimH-1 {
			continue
		}
		for cIdx, cell := range line {
			if cIdx == 0 || cIdx == dimW-1 {
				continue
			}
			if cell.letter == 3 {
				isXMAS := checkA(rIdx, cIdx, cells)
				if isXMAS {
					total += 1
				}
			}
		}
	}
	return total
}
