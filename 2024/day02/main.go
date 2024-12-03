package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
	"strings"
)

func parse(filePath string) ([][]int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    var reports [][]int
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		var levels []int
		for _, field := range fields {
			val, err := strconv.Atoi(field)
			if err != nil {
				return reports, err
			}
			levels = append(levels, val)
		}
		reports = append(reports, levels)
	}

	return reports, nil
}

func isSafeRange(x, y int) bool {
	if x > y {
		return (x - y) >= 1 && (x - y) <= 3
	}
	return (y - x) >= 1 && (y - x) <= 3
}

func removeAtIndex(levels []int, i int) []int {
	tmp := make([]int, len(levels) - 1)
	copy(tmp, levels[:i])
	copy(tmp[i:], levels[i+1:])
	return tmp
}

func isSafeLevels(levels []int) bool {
	var isInc = levels[len(levels) - 1] > levels[0]
	var isSafe = true
	for i, level := range levels {
		if i == 0 {
			continue
		}
		prev := levels[i - 1]
		if isInc && prev >= level {
			isSafe = false
			break
		}
		if !isInc && prev <= level {
			isSafe = false
			break
		}

		if !isSafeRange(prev, level) {
			isSafe = false
			break
		}
	}
	return isSafe
}

func part1(reports [][]int) {
	var count = 0
	for _, levels := range reports {
		isSafe := isSafeLevels(levels)
		if isSafe {
			count++
		}
	}
	fmt.Println(count)
}

func part2(reports [][]int) {
	var count = 0
	for _, levels := range reports {
		isSafe := isSafeLevels(levels)
		if isSafe {
			count++
		} else {
			// try again by removing an item
			for i := 0; i < len(levels); i++ {
				tmp := removeAtIndex(levels, i)
				if (isSafeLevels(tmp)) {
					count++
					break
				}
			}
		}
	}
	fmt.Println(count)
}


func main() {
	reports, err := parse("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	part1(reports)
	part2(reports)
}
