package day06

import (
	"AoC_2024/util"
	"errors"
	"strings"
)

// corresponds to each direction (eg cul = const up left)
const (
	cuu = iota + 1
	crr
	cdd
	cll
)

// var dirs = []uint8{0, 1, 2, 3}

func moveCellIndex(dir uint8, rIdx int, cIdx int) (newRIdx int, newCIdx int, err error) {
	switch dir {
	case cuu:
		newRIdx = rIdx - 1
		newCIdx = cIdx
	case cll:
		newRIdx = rIdx
		newCIdx = cIdx - 1
	case crr:
		newRIdx = rIdx
		newCIdx = cIdx + 1
	case cdd:
		newRIdx = rIdx + 1
		newCIdx = cIdx
	}
	if newRIdx < 0 || newCIdx < 0 || newRIdx >= dimH || newCIdx >= dimW {
		err = errors.New("out of bounds")
		newCIdx = cIdx
		newRIdx = rIdx
	} else {
		err = nil
	}
	return
}

type Grid [dimH][dimW]uint8

var runeToUint = map[rune]uint8{
	rune('.'): 0,
	rune('#'): 1,
	rune('^'): 2,
	rune('X'): 3, // won't be in input
}

func Day06Main() {
	var grid Grid
	begR := 0
	begC := 0
	for r, line := range strings.Split(puzzleInput, "\n") {
		for c, char := range line {
			if runeToUint[char] == 2 {
				begR = r
				begC = c
			}
			grid[r][c] = runeToUint[char]
		}
	}
	util.ResultString(6, d06p01(grid, begR, begC), d06p02(grid, begR, begC))
}

const puzzleInput = d06input

// const dimH = 10
// const dimW = 10

const dimH = 130
const dimW = 130

const d06testinput = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

const d06input = `...........#....................#......#.....................#.#...........#......................................................
.....................................#....#.................#...............................#.....................................
..#..#.......#.............................................#.#.........................#......................#..............#....
.........................#.....#...............#.................#................................................................
.......#..........#........#.#............................................#....#.....................#..........##..#.............
.........#.....................................................#..............................................................#...
.....#..............................................................#................##.....#...........#.......#.........#.......
.......#.........#......#...............#............................#.........#.....#.....#...#..................................
.......#...............#...................#...............#............#.......................................................#.
........#.......................................................................##..#..............#...#..........................
..............#........#.......................#......#...............#..#......................#......#...#............#.........
................#................#...........................................#.....#................#..........#..................
..##....................#...........#......................#..................#.......#...........................................
..................................................................#..........#...................#................................
........................#.......#..........................#................................#.....................#...#.#.........
...............#........#...............#.##..#.#...............#...........#......................................#..............
.#.........#.#...............................................................................#..........................#...#.....
..............................#.....#.............##..............#.##....................#......#.................#....#.........
..............#.............................#........#.....#..........................................#...........#...............
...................#.................#..............#....##..#......................#..............................#....#.#.......
.............##.....................#..........#....................................#.....................#...#.#.................
.........................#..........#.......#.................................#...................................................
............#.................#.................................................#.........##....................#.........#.....#.
.#.........#.........................#..............#............#..#...........................##.........#......................
...........................#.....#...................#....#...........#.#.#...........#...................#.......................
.................#..#......#..............................................................................#...................#...
...............................#..............#...#..#......##.#.........................#........#...........#...................
........#...............##.............................................................#..........................#...............
.............#.............#.............#...................#...........#....................#...................................
##.......................................................................................#....#...........#....#...........#......
.....................#.#.......................................#.....................................#.................#..........
......................#...................................................................#.....................#............#...#
.........................#.#........................#.#................#....#....#.....#..........................................
.....................#...#.......................................#.......#.............#...................#......................
.............#...............#...............................................................................................#....
........................................................................................................#........#.........#......
.................................##....................#....#..............#.....#.....#..#....#.........................#.#.....#
.........................................................................#........#........#.........................#...#........
#...................#..........#......................#..............................#.#........................#.................
.#.....................................................................#..#.......#........................#......................
.......................#..............#..............#..#..............................#..........................................
...................................................................................................#...................#..........
...............#............#....................................#.....................#....................#.....................
............#..#.#..........#.#.......#..........................#.................#.#..........#.......#.........................
.......#....#.............................#..............................#..#........................#........#...................
........#.................#.....#.#..............................#........................#...................#..........#........
............................#........................#.........................................................#..................
....................................#.............##....................................................................#..#......
..................................................#................................................#..............#........#...##.
...........##.................#.........................#....................................................#...........#........
...........#.#..................#.......................#...#.......#................#.......##...................................
................................................#...........#........#.......................#.....................#....#.#.......
...............#........................................#............................#.....................................#......
........#...#.....#.......................................................#.#................................#.#..................
.....#............#.#...............#.#.....#....................#.#...................#...................##.........#..........#
............#.................#....................................................#...........................................#..
................#..#.....#....................................#................#...#..............................................
.....................................................................#...............#............................................
..................#...........#.....#...................#....#..............................................#.....................
.#............#...................................#.........#.........................................#...........#...............
#.......#..#...#.....#............................................................................................................
..#.........................................................##..............................#................................#....
..............................#..............#........................................................#....................#......
#.........#.........#.............#..............#............................#..........#........................#..#............
....#...........#..........................##......#.............................#.....#........................#.................
.....#.................................................................#..#.......................................................
..........................#.......#.......#..................#....................................................................
................#....................#..................#.........................................................................
..............................................#..........................#...........#.....#...........................#....#.#...
.##...............#...#...........................................................................................................
..................#..#.......................#...#..#.#.................................#.........................#...........#...
..........................................................................#....................................#..................
#...........#.............................##.....................#........#....#...#........#.....................#.....#.........
.......#.............................#...^....#..................................#.........#......................................
........#..............................#................................#.....#....##.............................................
.................#...........................#........................#..........#..........#.................#.#..........#.....#
..................................................................................................................#...............
..........#.........#........................#........................................................#........................#..
................................................................................................................#.................
....................#......................#....#.......#.....#...................................................................
................................#........................................................#.........................###............
......#..............#.#....................#............#.......................#.......................................#........
.........................#....................................#..............#.......#............................................
..............#.................#........................................................#..#.....................................
.....#.#........................#...........................................#..#.....#.................#................#......##.
................#.............#...................................................................................................
#....................#..........#..............##...#........#..#.........................#............#......#..................#
...........#..#.......#.....................#..........................................................#........#.................
...............................................#................#...........#................................#....................
.............#.......#.................#........#.............................................#...........#.......................
....................#.......#.....#.......................#.................#.............#.........................#.............
.#..............#.................................................................................................................
..#.#......................................#.................#.........#...................................#.#....................
......#..............................#........#.#..................#................#.................................#...........
..............#........................................#..........................................................................
.#...........#.........#...........#...................#................#..................................#............#..#......
.........#....#...................#..#.................................................................#....#.....................
..........#...........................#.....................................#...............#.......#...#.........................
...........#..................................#............................................................##................#....
.......................#..#.......................................................................................................
...............................................#......................#.#.#.............#..#........#...........................#.
..#.................................................#..............................#.....................................#........
..........#..#..........#...................#......................................##......#.......#........#.....#...............
.#.#...............................................##...#................#..............#.....#...................................
....................#..#...............................#..#.......#....................#....#.....................#.....#.#.......
#...................#...........#...................................................#...........................#..##.........###.
.............#.........#......#..........................#.................#......##................#...............#...#.........
........#.....................................................................................#...................................
..#...........#...................................#...................#............................#.........#.................#..
.................................................................................#................................................
.#........#.................#......................................#.....#...........#................#...........................
#...........................#..............................#....#............................................................#....
#..........#.............................................#..........#.......................................................#.....
.........#...#...#...#............................................................#....#..........#.......................#.......
.........................#.........................#........#............#..............#.....##.........................#...#....
...............................#...........#..............................................#......................#......#.........
......................#........................#..#.....................#...............................#........................#
......#...##..............................................................................................................#.......
............#.............#.#...............................................#.........#..........................#...........#....
.......................................#............#...............................................#.............................
....#........#.............................#...............#.#................................................#...................
.........#.#........................#.....#...##...#.........###............#.#.........#......#.........#................#.....#.
...........#....................#......................##...............#.......................................................#.
#......................................................#.................##....................................#........#.....#...
..#..........................##.................#...........................................#.....................................
.......................................#............#......#...#.........#......##......................................#.........
.#.......#.........................................##..............................................#.#...................#........
.................#.........#..........................##............................#..#.....#....................................
.....#....................................#......................................#.......................#...........#..#..#......
.......#......#.........#.....#........#......#.................#...........#............#..................#....#...........#....`
