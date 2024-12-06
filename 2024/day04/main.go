package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Pos struct {
	row int
	col int
}

// XMAS
var XMAS = "XMAS"

func dfsRec(rows [][]string, pos Pos, direction string, index int) int {
	if index >= len(XMAS) {
		return 1
	}

	if pos.row < 0 || pos.row >= len(rows) || pos.col < 0 || pos.col >= len(rows[0]) {
		return 0
	}

	if rows[pos.row][pos.col] != string(XMAS[index]) {
		return 0
	}

	index++

	switch direction {
	case "T":
		return dfsRec(rows, Pos{pos.row - 1, pos.col}, direction, index)
	case "TR":
		return dfsRec(rows, Pos{pos.row - 1, pos.col + 1}, direction, index)
	case "R":
		return dfsRec(rows, Pos{pos.row, pos.col + 1}, direction, index)
	case "BR":
		return dfsRec(rows, Pos{pos.row + 1, pos.col + 1}, direction, index)
	case "B":
		return dfsRec(rows, Pos{pos.row + 1, pos.col}, direction, index)
	case "BL":
		return dfsRec(rows, Pos{pos.row + 1, pos.col - 1}, direction, index)
	case "L":
		return dfsRec(rows, Pos{pos.row, pos.col - 1}, direction, index)
	case "TL":
		return dfsRec(rows, Pos{pos.row - 1, pos.col - 1}, direction, index)
	}

	return 0
}

func dfs(rows [][]string, start Pos) int {
	count := 0
	directions := []string{
		"T", "TR", "R", "BR", "B", "BL", "L", "TL",
	}
	for _, d := range directions {
		count += dfsRec(rows, start, d, 0)
	}

	return count
}

func part1(rows [][]string) {
	count := 0
	for r, row := range rows {
		for c := range row {
			res := dfs(rows, Pos{r, c})
			count += res
		}
	}
	fmt.Println(count)
}

var MAS = "MAS"

func checkA(rows [][]string, pos Pos) bool {
	if pos.row-1 < 0 || pos.row+1 >= len(rows) || pos.col-1 < 0 || pos.col+1 >= len(rows[0]) {
		return false
	}

	// BR MAS
	/*
		M.S
		.A.
		M.S
	*/
	if rows[pos.row-1][pos.col-1] == rows[pos.row+1][pos.col-1] &&
		rows[pos.row-1][pos.col-1] == "M" &&
		rows[pos.row-1][pos.col+1] == rows[pos.row+1][pos.col+1] &&
		rows[pos.row-1][pos.col+1] == "S" {
		return true
	}
	/*
		S.M
		.A.
		S.M
	*/

	if rows[pos.row-1][pos.col-1] == rows[pos.row+1][pos.col-1] &&
		rows[pos.row-1][pos.col-1] == "S" &&
		rows[pos.row-1][pos.col+1] == rows[pos.row+1][pos.col+1] &&
		rows[pos.row-1][pos.col+1] == "M" {
		return true
	}

	/*
		M.M
		.A.
		S.S
	*/
	if rows[pos.row-1][pos.col-1] == rows[pos.row-1][pos.col+1] &&

		rows[pos.row-1][pos.col-1] == "M" &&
		rows[pos.row+1][pos.col-1] == rows[pos.row+1][pos.col+1] &&

		rows[pos.row+1][pos.col-1] == "S" {
		return true
	}

	/*
		S.S
		.A.
		M.M
	*/
	if rows[pos.row-1][pos.col-1] == rows[pos.row-1][pos.col+1] &&

		rows[pos.row-1][pos.col-1] == "S" &&
		rows[pos.row+1][pos.col-1] == rows[pos.row+1][pos.col+1] &&

		rows[pos.row+1][pos.col-1] == "M" {
		return true
	}

	return false
}

func part2(rows [][]string) {
	count := 0
	for r, row := range rows {
		for c, col := range row {
			if col == "A" && checkA(rows, Pos{r, c}) {
				count++
			}
		}
	}
	fmt.Println(count)
}

func parse(filePath string) [][]string {
	var rows [][]string
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var chars []string
		for _, c := range line {
			chars = append(chars, string(c))
		}
		rows = append(rows, chars)
	}
	return rows
}

func main() {
	rows := parse("./input.txt")
	//part1(rows)
	part2(rows)
}
