package day04

import "errors"

func moveCellIndex(dir uint8, rIdx int, cIdx int) (newRIdx int, newCIdx int, err error) {
	switch dir {
	case cul:
		newRIdx = rIdx - 1
		newCIdx = cIdx - 1
	case cuu:
		newRIdx = rIdx - 1
		newCIdx = cIdx
	case cur:
		newRIdx = rIdx - 1
		newCIdx = cIdx + 1
	case cll:
		newRIdx = rIdx
		newCIdx = cIdx - 1
	case crr:
		newRIdx = rIdx
		newCIdx = cIdx + 1
	case cdl:
		newRIdx = rIdx + 1
		newCIdx = cIdx - 1
	case cdd:
		newRIdx = rIdx + 1
		newCIdx = cIdx
	case cdr:
		newRIdx = rIdx + 1
		newCIdx = cIdx + 1
	}
	if newRIdx < 0 || newCIdx < 0 || newRIdx >= dimH || newCIdx >= dimW {
		err = errors.New("Out of Bounds")
		newCIdx = cIdx
		newRIdx = rIdx
	} else {
		err = nil
	}
	return
}

func genPresList(cells Cells, rIdx int, cIdx int) (newPresList PresList, noYes int) {
	newPresList = initialPres
	thisCell := cells[rIdx][cIdx]
	noYes = 0 // number of ye's
	for idx, dir := range dirs {
		newRIdx, newCIdx, err := moveCellIndex(dir, rIdx, cIdx)
		if err != nil {
			continue
		}
		adjCell := cells[newRIdx][newCIdx]
		if adjCell.letter == thisCell.letter-1 {
			if adjCell.pres[idx] == ye {
				noYes += 1
			}
			newPresList[idx] = adjCell.pres[idx]
		}
	}
	return
}

// kinda overcomplicated this one
func d04p01(cells Cells) int {
	total := 0
	for _, letter := range [4]uint8{2, 3, 4} {
		for rIdx, cellRow := range cells {
			for cIdx, cell := range cellRow {
				if cell.letter == letter {
					switch cell.letter {
					case 2, 3:
						newCell := cell
						newCell.pres, _ = genPresList(cells, rIdx, cIdx)
						cells[rIdx][cIdx] = newCell
					case 4:
						_, noYes := genPresList(cells, rIdx, cIdx)
						total += noYes
					}
				}
			}
		}
	}
	return total
}
