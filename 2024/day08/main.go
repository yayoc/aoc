package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func getDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func getAntinodes(r1, c1, r2, c2 int) [][]int {
	res := make([][]int, 0)
	rDiff := getDiff(r1, r2)
	cDiff := getDiff(c1, c2)

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	if r1 < r2 {
		a1 = append(a1, r1-rDiff)
		a2 = append(a2, r2+rDiff)
	} else {
		a1 = append(a1, r1+rDiff)
		a2 = append(a2, r2-rDiff)
	}

	if c1 < c2 {
		a1 = append(a1, c1-cDiff)
		a2 = append(a2, c2+cDiff)
	} else {
		a1 = append(a1, c1+cDiff)
		a2 = append(a2, c2-cDiff)
	}

	res = append(res, a1)
	res = append(res, a2)
	return res
}

func getAllAntiNodes(maxR, maxC, r1, c1, r2, c2 int) [][]int {
	res := make([][]int, 0)
	rDiff := getDiff(r1, r2)
	cDiff := getDiff(c1, c2)

	// extract minimum diff
	if rDiff > cDiff && rDiff%cDiff == 0 {
		rDiff = rDiff / cDiff
		cDiff = cDiff / cDiff
	} else if cDiff > rDiff && cDiff%rDiff == 0 {
		cDiff = cDiff / rDiff
		rDiff = rDiff / rDiff
	}

	a1 := make([]int, 0)
	a2 := make([]int, 0)

	a1rDiff := 0
	a1cDiff := 0
	a2rDiff := 0
	a2cDiff := 0

	if r1 < r2 {
		a1rDiff = r1 - rDiff
		a2rDiff = r2 + rDiff
	} else {
		a1rDiff = r1 + rDiff
		a2rDiff = r2 - rDiff
	}

	if c1 < c2 {
		a1cDiff = c1 - cDiff
		a2cDiff = c2 + cDiff
	} else {
		a1cDiff = c1 + cDiff
		a2cDiff = c2 - cDiff
	}
	a1 = append(a1, a1rDiff)
	a1 = append(a1, a1cDiff)

	a2 = append(a2, a2rDiff)
	a2 = append(a2, a2cDiff)

	for a1[0] >= 0 && a1[0] <= maxR && a1[1] >= 0 && a1[1] <= maxC {
		tmp := make([]int, len(a1))
		copy(tmp, a1)
		res = append(res, tmp)
		if r1 < r2 {
			a1[0] -= rDiff
		} else {
			a1[0] += rDiff
		}
		if c1 < c2 {
			a1[1] -= cDiff
		} else {
			a1[1] += cDiff
		}
	}

	for a2[0] >= 0 && a2[0] <= maxR && a2[1] >= 0 && a2[1] <= maxC {
		tmp := make([]int, len(a2))
		copy(tmp, a2)
		res = append(res, tmp)
		if r1 < r2 {
			a2[0] += rDiff
		} else {
			a2[0] -= rDiff
		}
		if c1 < c2 {
			a2[1] += cDiff
		} else {
			a2[1] -= cDiff
		}
	}

	return res
}

/*
{1,2,3}
=> {{1,2}, {2,3}, {1,3}}
*/

func pickAllPairs(nums [][]int) [][][]int {
	var pairs [][][]int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			pair := make([][]int, 0)
			pair = append(pair, nums[i])
			pair = append(pair, nums[j])
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func makeMap(grid [][]string) map[string][][]int {
	m := make(map[string][][]int)
	for ri, row := range grid {
		for ci, col := range row {
			_, ok := m[col]
			if !ok {
				m[col] = make([][]int, 0)
			}
			m[col] = append(m[col], []int{ri, ci})
		}
	}

	return m
}

func part1(grid [][]string) {
	// keep antenna locations
	m := makeMap(grid)

	visited := make(map[[2]int]bool)
	for key, val := range m {
		if key != "." && len(val) > 2 {
			// pick two elements
			pairs := pickAllPairs(val)
			for _, pair := range pairs {
				// get anti nodes
				nodes := getAntinodes(pair[0][0], pair[0][1], pair[1][0], pair[1][1])
				// place them if possible
				for _, node := range nodes {
					r := node[0]
					c := node[1]
					// check boundary
					if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) {
						continue
					}

					if !visited[[2]int{r, c}] {
						visited[[2]int{r, c}] = true
					}

					if grid[r][c] == "." {
						grid[r][c] = "#"
					}
				}
			}
		}
	}
	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(len(visited))
}

func part2(grid [][]string) {
	m := makeMap(grid)
	visited := make(map[[2]int]bool)

	for key, val := range m {
		if key != "." && len(val) > 2 {
			// pick two elements
			pairs := pickAllPairs(val)
			for _, pair := range pairs {
				// get anti nodes
				nodes := getAllAntiNodes(len(grid)-1, len(grid[0])-1, pair[0][0], pair[0][1], pair[1][0], pair[1][1])

				for _, p := range pair {
					row := p[0]
					col := p[1]
					if !visited[[2]int{row, col}] {
						visited[[2]int{row, col}] = true
					}
				}

				// place them if possible
				fmt.Println("anti nodes:", nodes)
				for _, node := range nodes {
					r := node[0]
					c := node[1]
					if !visited[[2]int{r, c}] {
						visited[[2]int{r, c}] = true
					}
					if grid[r][c] == "." {
						grid[r][c] = "#"
					}
				}
			}
		}
	}

	for _, row := range grid {
		fmt.Println(row)
	}
	fmt.Println(len(visited))
}

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

func main() {
	grid := parse("./input.txt")
	// part1(grid)
	part2(grid)
}
