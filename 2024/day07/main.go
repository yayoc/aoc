package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"fmt"
)

var OPERATORS = []string{"+", "*"}

func parse(filePath string) map[int][]int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		result := strings.Split(line, ":")
		left, _ := strconv.Atoi(result[0])
		_, ok := res[left]
		if !ok {
			res[left] = make([]int, 0)
		}
		v := strings.Fields(result[1])
		for _, vv := range v {
			vvv, _ := strconv.Atoi(vv)
			res[left] = append(res[left], vvv)
		}
	}
	return res
}

func calc(nums []int, ops []string) int {
	for len(nums) > 0 {
		if len(nums) == 1 {
			return nums[0]
		}

		l := nums[0]
		r := nums[1]
		op := ops[0]
		var tmp int
		switch op {
		case "+":
			tmp += (l + r)
		case "*":
			tmp += (l * r)
		case "||":
			res := strconv.Itoa(l) + strconv.Itoa(r)
			n, _ := strconv.Atoi(res)
			tmp += n
		}
		ops = ops[1:]
		nums = append([]int{tmp}, nums[2:]...)
	}
	return nums[0]
}

func permutateOps(ops []string, numAdd int) [][]string {
	var result [][]string
	l := len(ops)

	if numAdd > l {
		return result
	}

	var generate func(index, plusLeft int)
	curr := append([]string{}, ops...)

	generate = func(index, plusLeft int) {
		if plusLeft == 0 {
			result = append(result, append([]string{}, curr...))
			return 
		}

		if index == l {
			return
		}

		original := curr[index]
		curr[index] = "+"
		generate(index+1, plusLeft-1)
		curr[index] = original

		generate(index+1, plusLeft)
	}

	generate(0, numAdd)
	return result
}

func part1(m map[int][]int) {
	count := 0
	for k, v := range m {
		l := len(v)
		ops := make([]string, l - 1)
		for i := range ops {
			ops[i] = "*"
		}
		numAdd := 0
		out:
		for numAdd <= l - 1 {
			permutations := permutateOps(ops, numAdd)
			fmt.Println(permutations)
			for _, p := range permutations {
				res := calc(v, p)
				if k == res {
					// found
					count += k
					break out
				}
			}
			numAdd++
		}
	}
	fmt.Println(count)
}

func permutations(elements []string, length int) [][]string {
	if length == 0 {
		return [][]string{{}}
	}

	var result [][]string
	subPerms := permutations(elements, length-1) // All permutations of length-1
	for _, sub := range subPerms {
		for _, e := range elements {
			newCombination := append([]string(nil), sub...)
			newCombination = append(newCombination, e)
			result = append(result, newCombination)
		}
	}
	return result
}

func part2(m map[int][]int) {
	ops := []string{"*", "+", "||"}
	count := 0
	for k, v := range m {
		l := len(v)
		permutations := permutations(ops, l - 1)
		out:
		for _, p := range permutations {
			res := calc(v, p)
			if k == res {
				// found
				fmt.Println(res, v, p)
				count += k
				break out
			}
		}
	}
	fmt.Println(count)
}

func main() {
	m := parse("./input.txt")
	// part1(m)
	part2(m)
}
