package main

import (
    "os"
	"log"
	"bufio"
	"strings"
	"fmt"
)


func parse(filePath string) [][]string {
    file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, "")
		grid = append(grid, v)
	}
	return grid
}

func bfs(grid [][]string, start [2]int) [][2]int {
	R := len(grid)
	C := len(grid[0])
	queue := make([][2]int, 0)
	queue = append(queue, start)
	visited := make(map[[2]int]bool)
	blocks := make([][2]int, 0)

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		row := q[0]
		col := q[1]

		if visited[[2]int{row, col}] {
			continue
		}

		visited[q] = true
		blocks = append(blocks, q)
		dirs := [][2]int{{-1,0}, {0,1}, {1,0}, {0,-1}}

		for _, dir := range dirs {
			newRow := row + dir[0]
			newCol := col + dir[1]

			if newRow >= 0 && newRow < R && newCol >= 0 && newCol < C {
				if grid[newRow][newCol] == grid[start[0]][start[1]] {
					queue = append(queue, [2]int{newRow, newCol})
				}
			}
		}
	}
	return blocks
}

func has(blocks [][2]int, r, c int) bool {
	for _, block := range blocks {
		row := block[0]
		col := block[1]
		if row == r && col == c {
			return true
		}
	}
	return false
}

func count(grid [][]string, originalRow, originalCol int) int {
	res := 0
	dirs := [][2]int{{-1,0}, {0,1}, {1,0}, {0,-1}}
	for _, dir := range dirs {
		row := originalRow + dir[0]
		col := originalCol + dir[1]
		if row < 0 || row >= len(grid) || col < 0 || col >= len(grid[0]) {
			res++
			continue
		}
		if grid[originalRow][originalCol] != grid[row][col] {
			res++
		}
	}
	return res
}

func minmax(blocks [][2]int) (int, int, int, int) {
	minRow := MaxInt
	maxRow := 0
	minCol := MaxInt
	maxCol := 0

	for block := range blocks {
		row := block[0]
		col := block[1]

		if row < minRow {
			minRow = row
		}
		if row > maxRow {
			maxRow = row
		}
		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}
	}

	return minRow, maxRow, minCol, maxCol
}


// calc a number of corners
func calcSides(grid [][]string, blocks [][2]int) int {
	if len(blocks) == 0 {
		return 0
	}
	corners := 0
	dirs := [][2]int{{-1,0}, {0,1}, {1,0}, {0,-1}}
	minRow, maxRow, minCol, maxCol = minmax(blocks)
	val := grid[blocks[0][0]][blocks[0][1]]
	for row := minRow; row <= maxRow; row++ {
		for col := minCol; col <= maxCol; col++ {
			if grid[row][col] == val {
				// check outer corner
			} else {
				// check inner corner
			}
		}
	}
	return corners
}

func part1(grid [][]string) {
	blocks := make([][2]int, 0)
	res := 0
	for r := range grid {
		for c := range grid[r] {
			if !has(blocks, r, c) {
				fmt.Println(grid[r][c])
				newBlocks := bfs(grid, [2]int{r,c})
				fmt.Println("newBlocks", newBlocks)
				// count fences
				for _, block := range newBlocks {
					newC := count(grid, block[0], block[1])
					fmt.Println(block)
					fmt.Println(newC)
					res += len(newBlocks) * newC
				}

				blocks = append(blocks, newBlocks...)
			}
		}
	}
	fmt.Println(res)
}

func main() {
	grid := parse("./input.txt")
	part1(grid)
}
