package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func absDiff(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var left []int
	var right []int
	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Fields(line)
		i, err := strconv.Atoi(result[0])
		if err != nil {
			log.Fatal(err)
		}
		left = append(left, i)
		ii, err := strconv.Atoi(result[1])
		if err != nil {
			log.Fatal(err)
		}
		right = append(right, ii)
	}

	part1 := func() {
		sort.Ints(left)
		sort.Ints(right)

		var res = 0
		for i, v := range left {
			res += absDiff(v, right[i])
		}
		fmt.Println(res)
	}

	part2 := func() {
        m := make(map[int]int)
		for _, v := range left {
			m[v] = 0
		}

		for _, v := range right {
			_, ok := m[v]
			if ok {
				m[v]++
			}
		}
		var res = 0
		for key, value := range m {
			res += (key * value)
		}
		fmt.Println(res)
	}

	part1()
    part2()

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
