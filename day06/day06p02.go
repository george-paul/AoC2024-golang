package day06

import (
	// "fmt"
	"sync"
	"sync/atomic"
)

type DirLineCell [5]bool // 0th will always be
type DirLines [dimH][dimW]DirLineCell

// func markDirLines(dirLines *DirLines, grid Grid, startR int, startC int, blockDir uint8) {
// 	r := startR
// 	c := startC
// 	oppDir := ((blockDir + 1) % 4) + 1 // opp dir
// 	for {
// 		(*dirLines)[r][c][blockDir] = true
// 		newR, newC, err := moveCellIndex(oppDir, r, c)
// 		next := grid[newR][newC]
// 		if err != nil || next == 1 {
// 			break
// 		} else {
// 			r = newR
// 			c = newC
// 		}

// 	}
// }

// func dirLinesForBlock(dirLines *DirLines, grid Grid, blockR int, blockC int) {
// 	for _, dir := range []uint8{1, 2, 3, 4} {
// 		nextR, nextC, err := moveCellIndex(dir, blockR, blockC)
// 		if err == nil {
// 			next := grid[nextR][nextC]
// 			if next != 1 {
// 				oppDir := ((dir + 1) % 4) + 1 // opp dir
// 				markDirLines(dirLines, grid, nextR, nextC, oppDir)
// 			}
// 		}
// 	}
// }

// func d06p02(grid Grid, begR int, begC int) int {
// 	var dirLines DirLines
// 	// prep dirLines
// 	for r, row := range grid {
// 		for c, cell := range row {
// 			if cell == 1 {
// 				// is block
// 				dirLinesForBlock(&dirLines, grid, r, c)
// 			}
// 		}
// 	}

// 	total := 0
// 	r := begR
// 	c := begC
// 	curDir := uint8(cuu) // initial direction is up
// 	for {
// 		// check if on dirLine
// 		nextDir := (curDir % 4) + 1
// 		pastDirCell := dirLines[r][c]
// 		if pastDirCell[nextDir] {
// 			total++
// 		}
// 		newR, newC, err := moveCellIndex(curDir, r, c)
// 		next := grid[newR][newC]
// 		if err != nil {
// 			break
// 		}
// 		if next == 1 {
// 			// markDirLines(&dirLines, grid, r, c, curDir)
// 			curDir = nextDir
// 		} else {
// 			r = newR
// 			c = newC
// 		}
// 	}
// 	return total
// }

const maxCyclePeriod = dimH * dimW * 10

func checkObstacleWorker(grid Grid, obsR int, obsC int, guardDir uint8, total *uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	// defer fmt.Println(obsR, ",", obsC, " done")
	// fmt.Println(obsR, ",", obsC)
	stepC := 0
	grid[obsR][obsC] = 1
	oppDir := ((guardDir + 1) % 4) + 1 // opp dir
	r, c, _ := moveCellIndex(oppDir, obsR, obsC)
	curDir := guardDir
	for {
		newR, newC, err := moveCellIndex(curDir, r, c)
		next := grid[newR][newC]
		if err != nil {
			return
		}
		if next == 1 {
			curDir = (curDir % 4) + 1 // turn right
			stepC += 1
		} else {
			r = newR
			c = newC
			stepC += 1
			// fmt.Println(obsR, ",", obsC, "; step:", stepC)
			if stepC >= maxCyclePeriod {
				atomic.AddUint32(total, 1)
				return
			}
		}
	}
}

func d06p02(grid Grid, begR int, begC int) int {
	r := begR
	c := begC
	curDir := uint8(cuu) // initial direction is up
	// total := make(chan int)
	total := uint32(0)
	var wg sync.WaitGroup

	for {
		newR, newC, err := moveCellIndex(curDir, r, c)
		next := grid[newR][newC]
		if err != nil {
			break
		}
		if next == 1 {
			curDir = (curDir % 4) + 1 // turn right
		} else {
			r = newR
			c = newC
			wg.Add(1)
			go checkObstacleWorker(grid, r, c, curDir, &total, &wg)
		}
	}
	wg.Wait()

	// sum increments
	counter := 0
	for range total {
		counter++
	}
	return counter
}
