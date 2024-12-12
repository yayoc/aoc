package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func parse(filePath string) []string {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	chars := []rune(str)
	var result []string
	for _, r := range chars {
		if !unicode.IsSpace(r) {
			result = append(result, string(r))
		}
	}
	return result
}

func appendMultiple(slice []string, elem string, count int) []string {
	for i := 0; i < count; i++ {
		slice = append(slice, elem)
	}
	return slice
}

func shift(blocks []string) []string {
	l := 0
	r := len(blocks) - 1

	for l < r {

		for l < r && blocks[l] != "." {
			l++
		}

		for l < r && blocks[r] == "." {
			r--
		}

		blocks[l] = blocks[r]
		blocks[r] = "."
		l++
		r--
	}

	return blocks
}

func find(blocks []string, space int, r int) int {
	l := 0
	for l < r {

		for l < r && blocks[l] != "." {
			l++
		}

		tmp := l
		for l < r {
			if blocks[l] != "." {
				break
			}
			if l-tmp >= space-1 {
				return tmp
			}
			l++
		}
	}

	return -1
}

func updateRange(blocks []string, start, end int, value string) []string {
	for i := start; i < end; i++ {
		blocks[i] = value
	}
	return blocks
}

func shift2(blocks []string) []string {
	counter := make(map[string]int)
	for _, block := range blocks {
		if block != "." {
			counter[block]++
		}
	}
	l := 0
	r := len(blocks) - 1

	for l < r {
		for l < r && blocks[r] == "." {
			r--
		}

		blockLen := counter[blocks[r]]
		index := find(blocks, blockLen, r)
		if index > 0 {
			blocks = updateRange(blocks, index, index+blockLen, blocks[r])
			blocks = updateRange(blocks, r-blockLen+1, r+1, ".")
		}
		r = r - blockLen
	}
	return blocks
}

func checksum(blocks []string) int {
	res := 0
	for i, block := range blocks {
		if block != "." {
			n, _ := strconv.Atoi(block)
			res += i * n
		}
	}
	return res
}

func getBlocks(chars []string) []string {
	var blocks []string
	id := 0
	i := 0
	for i < len(chars) {
		num, _ := strconv.Atoi(chars[i])
		if i%2 == 0 {
			blocks = appendMultiple(blocks, strconv.Itoa(id), num)
			id++
		} else {
			blocks = appendMultiple(blocks, ".", num)
		}
		i++
	}
	return blocks
}

func part1(chars []string) {
	blocks := getBlocks(chars)
	blocks = shift(blocks)
	fmt.Println(checksum(blocks))
}

func part2(chars []string) {
	blocks := getBlocks(chars)
	fmt.Println(blocks)
	blocks = shift2(blocks)
	fmt.Println(blocks)
	fmt.Println(checksum(blocks))
}

func main() {
	chars := parse("./input.txt")
	//part1(chars)
	part2(chars)
}
