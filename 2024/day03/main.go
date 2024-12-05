package main

import (
	"log"
	"os"
	"strconv"
	"fmt"
    "regexp"
)

func isNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

type Mul struct {
    left int
    right int
}

type Do struct {
    start int
    end int
}

type Dont struct {
    start int
    end int
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

func part1(muls []Mul) {
    sum := 0
    for _, m := range muls {
        sum += m.left * m.right 
    }
    fmt.Println(sum)
}

func part2(muls []Mul) {
}

func main() {
    muls := parse("./input.txt")
    part1(muls)
}
