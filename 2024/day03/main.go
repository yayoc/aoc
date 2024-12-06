package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Mul struct {
	left  int
	right int
}

func parse(filePath string) []Mul {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	var muls []Mul

	re, err := regexp.Compile(`mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		left, _ := strconv.Atoi(match[re.SubexpIndex("left")])
		right, _ := strconv.Atoi(match[re.SubexpIndex("right")])
		muls = append(muls, Mul{left, right})
	}

	return muls
}

func isDisabled(str string) bool {
	re, err := regexp.Compile(`(?P<do>do\(\))|(?P<dont>don't\(\))`)
	if err != nil {
		log.Fatal(err)
	}
	matches := re.FindAllStringSubmatchIndex(str, -1)
	if len(matches) == 0 {
		return false
	}
	last := matches[len(matches)-1]
	dont := last[re.SubexpIndex("dont")*2]
	if dont != -1 {
		return true
	}
	return false
}

func parse2(filePath string) []Mul {
	b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
	var muls []Mul

	re, err := regexp.Compile(`mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\)`)
	if err != nil {
		log.Fatal(err)
	}

	matches := re.FindAllStringSubmatchIndex(str, -1)

	for _, match := range matches {
		start := match[0]
		end := match[1]
		prev_str := str[:start]
		if !isDisabled(prev_str) {
			ms := re.FindAllStringSubmatch(str[start:end], -1)
			for _, m := range ms {
				left, _ := strconv.Atoi(m[re.SubexpIndex("left")])
				right, _ := strconv.Atoi(m[re.SubexpIndex("right")])
				muls = append(muls, Mul{left, right})
			}
		}
	}

	return muls
}

func part1(muls []Mul) {
	sum := 0
	for _, m := range muls {
		sum += m.left * m.right
	}
	fmt.Println(sum)
}

func part2(muls []Mul) {
	sum := 0
	for _, m := range muls {
		sum += m.left * m.right
	}
	fmt.Println(sum)
}

func main() {
	muls := parse("./input.txt")
	muls2 := parse2("./input.txt")
	part1(muls)
	part2(muls2)
}
