package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parse(filePath string) [][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	grid := make([][]int, 0)
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, "")
		row := make([]int, 0)
		for _, vv := range v {
			i, _ := strconv.Atoi(vv)
			row = append(row, i)
		}
		grid = append(grid, row)
	}
	return grid

}

func getUniqueSize(ends [][2]int) int {
	m := make(map[[2]int]bool, 0)

	for _, e := range ends {
		m[[2]int{e[0], e[1]}] = true
	}

	return len(m)
}

func dfs(grid [][]int, row, col int, isUnique bool) int {
	R := len(grid)
	C := len(grid[0])

	stack := make([][3]int, 0)
	stack = append(stack, [3]int{row, col, 0})
	ends := make([][2]int, 0)

	for len(stack) > 0 {
		l := len(stack)
		p := stack[l-1]
		stack = stack[:l-1]

		r := p[0]
		c := p[1]
		h := p[2]

		if h == 9 {
			ends = append(ends, [2]int{r, c})
		}

		// (-1, 0), (0, 1), (0, -1), (1, 0)
		for _, v := range [][2]int{{-1, 0}, {0, 1}, {0, -1}, {1, 0}} {
			nr := r + v[0]
			nc := c + v[1]
			nh := h + 1

			if nr >= 0 && nr < R && nc >= 0 && nc < C {
				if grid[nr][nc] == nh {
					stack = append(stack, [3]int{nr, nc, nh})
				}
			}
		}

	}

	if isUnique {
		return len(ends)
	}
	return getUniqueSize(ends)
}

func part1(grid [][]int) {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				res := dfs(grid, r, c, false)
				count += res
			}
		}
	}
	fmt.Println(count)
}

func part2(grid [][]int) {
	count := 0
	for r := range grid {
		for c := range grid[r] {
			if grid[r][c] == 0 {
				res := dfs(grid, r, c, true)
				count += res
			}
		}
	}
	fmt.Println(count)
}

func main() {
	grid := parse("./input.txt")
	part1(grid)
	part2(grid)
}
