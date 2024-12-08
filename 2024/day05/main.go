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

type Rule = map[int][]int
type Ops = [][]int

func parse(filePath string) (Rule, Ops) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	isOps := false
	rule := make(map[int][]int)
	ops := [][]int{}

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isOps = true
			continue
		}

		if isOps {
			// parse ops
			result := strings.Split(line, ",")
			nums := []int{}
			for _, n := range result {
				i, _ := strconv.Atoi(n)
				nums = append(nums, i)
			}
			ops = append(ops, nums)
		} else {
			// parse rules
			result := strings.Split(line, "|")
			left, _ := strconv.Atoi(result[0])
			right, _ := strconv.Atoi(result[1])

			_, ok := rule[left]
			if !ok {
				rule[left] = make([]int, 0)
			}
			rule[left] = append(rule[left], right)
		}
	}

	return rule, ops
}

func fixOrder(rule Rule, op []int) []int {
	sort.Slice(op, func(i, j int) bool {
		afterI, _ := rule[op[i]]
		afterJ, _ := rule[op[j]]

		// only extract elements in op
		filteredI := make([]int, 0)
		filteredJ := make([]int, 0)

		for _, i := range afterI {
			for _, n := range op {
				if i == n {
					filteredI = append(filteredI, n)
				}
			}
		}

		for _, i := range afterJ {
			for _, n := range op {
				if i == n {
					filteredJ = append(filteredJ, n)
				}
			}
		}

		return len(filteredI) > len(filteredJ)
	})

	return op
}

func part1(rule Rule, ops Ops) {
	count := 0
	for _, op := range ops {
		isValid := true
	out:
		for i, n := range op {
			after, ok := rule[n]
			if !ok {
				continue
			}

			for _, a := range after {
				for _, b := range op[:i] {
					if a == b {
						isValid = false
						break out
					}
				}
			}
		}
		if isValid {
			fmt.Println("op is valid", op)
			count += op[len(op)/2]
		} else {
			fmt.Println("op is invalid", op)
		}
	}
	fmt.Println(count)
}

func isValid(rule Rule, op []int) bool {
	valid := true
out:
	for i, num := range op {
		after, ok := rule[num]
		if !ok {
			continue
		}
		for _, a := range after {
			for _, b := range op[:i] {
				if a == b {
					valid = false
					break out
				}
			}
		}
	}

	return valid
}

func part2(rule Rule, ops Ops) {
	count := 0
	for _, op := range ops {
		valid := isValid(rule, op)
		if valid {
			// count += op[len(op) / 2]
		} else {
			fmt.Println("invalid", op)
			op = fixOrder(rule, op)
			count += op[len(op)/2]
		}
	}
	fmt.Println(count)
}

func main() {
	rule, ops := parse("./input.txt")
	// part1(rule, ops)
	part2(rule, ops)
}
