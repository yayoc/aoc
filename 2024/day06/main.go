package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Guard struct {
	row       int
	col       int
	direction string
}

func parse(filePath string) ([][]bool, Guard) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Determine dimensions
	rows := len(lines)
	cols := 0
	if rows > 0 {
		cols = len(lines[0])
	}

	grid := make([][]bool, rows)
	for i := range grid {
		grid[i] = make([]bool, cols)
	}

	var guard Guard
	for r, line := range lines {
		for c, ch := range line {
			switch ch {
			case '#':
				grid[r][c] = true
			case '^':
				guard = Guard{r, c, "T"}
			case '>':
				guard = Guard{r, c, "R"}
			case 'v':
				guard = Guard{r, c, "B"}
			case '<':
				guard = Guard{r, c, "L"}
			default:
				// just a free cell
			}
		}
	}

	return grid, guard
}

func turn(guard *Guard) {
	switch guard.direction {
	case "T":
		guard.direction = "R"
	case "R":
		guard.direction = "B"
	case "B":
		guard.direction = "L"
	case "L":
		guard.direction = "T"
	}
}

func getNextPosition(grid [][]bool, guard Guard) (int, int, bool) {
	rows := len(grid)
	cols := len(grid[0])

	nextRow, nextCol := guard.row, guard.col
	switch guard.direction {
	case "T":
		nextRow--
	case "R":
		nextCol++
	case "B":
		nextRow++
	case "L":
		nextCol--
	}

	if nextRow < 0 || nextRow >= rows || nextCol < 0 || nextCol >= cols {
		return nextRow, nextCol, false
	}
	return nextRow, nextCol, true
}

func move(grid [][]bool, guard *Guard) bool {
	nextRow, nextCol, ok := getNextPosition(grid, *guard)
	if !ok {
		// Out of bounds
		return false
	}

	// If next is obstacle, turn until you find a non-obstacle direction or out-of-bounds
	for ok && grid[nextRow][nextCol] {
		turn(guard)
		nextRow, nextCol, ok = getNextPosition(grid, *guard)
		if !ok {
			return false
		}
	}

	if !ok {
		return false
	}

	guard.row = nextRow
	guard.col = nextCol
	return true
}

func isVisited(visited map[[2]int]bool, row, col int) bool {
	return visited[[2]int{row, col}]
}

func part1(grid [][]bool, guard Guard) {
	visited := make(map[[2]int]bool)
	visited[[2]int{guard.row, guard.col}] = true

	for move(grid, &guard) {
		if !isVisited(visited, guard.row, guard.col) {
			visited[[2]int{guard.row, guard.col}] = true
		}
	}
	fmt.Println(len(visited))
}

func isLoop(grid [][]bool, guard Guard) bool {
	fastGuard := guard

	// Each iteration, fastGuard moves twice, guard moves once
	for move(grid, &fastGuard) && move(grid, &fastGuard) {
		move(grid, &guard)
		if guard.row == fastGuard.row && guard.col == fastGuard.col && guard.direction == fastGuard.direction {
			return true
		}
	}
	return false
}

func part2(grid [][]bool, guard Guard) {
	visited := make(map[[2]int]bool)
	visited[[2]int{guard.row, guard.col}] = true

	tempGuard := guard
	for move(grid, &tempGuard) {
		visited[[2]int{tempGuard.row, tempGuard.col}] = true
	}

	originalGuard := guard
	count := 0
	for cell := range visited {
		originalVal := grid[cell[0]][cell[1]]
		grid[cell[0]][cell[1]] = true

		if isLoop(grid, originalGuard) {
			count++
		}

		grid[cell[0]][cell[1]] = originalVal
	}
	fmt.Println(count)
}

func main() {
	grid, guard := parse("input.txt")
	part1(grid, guard)
	part2(grid, guard)
}
