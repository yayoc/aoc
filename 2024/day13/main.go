package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	x int
	y int
}

type Prize struct {
	x int
	y int
}

type Game struct {
	buttonA Button
	buttonB Button
	prize   Prize
}

func parse(filePath string) []Game {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var buttonA Button
	var buttonB Button
	var prize Prize
	tmp := make([]Game, 0)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Contains(line, "Button") {
			reButton, err := regexp.Compile(`Button (?P<aorb>A|B): X\+(?P<x>\d+), Y\+(?P<y>\d+)`)
			if err != nil {
				log.Fatal(err)
			}
			matches := reButton.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				x, _ := strconv.Atoi(match[reButton.SubexpIndex("x")])
				y, _ := strconv.Atoi(match[reButton.SubexpIndex("y")])
				aOrB := match[reButton.SubexpIndex("aorb")]
				if aOrB == "A" {
					buttonA = Button{x, y}
				} else {
					buttonB = Button{x, y}
				}
			}
		} else if strings.Contains(line, "Prize") {
			rePrize, err := regexp.Compile(`Prize: X=(?P<x>\d+), Y=(?P<y>\d+)`)
			if err != nil {
				log.Fatal(err)
			}
			matches := rePrize.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				x, _ := strconv.Atoi(match[rePrize.SubexpIndex("x")])
				y, _ := strconv.Atoi(match[rePrize.SubexpIndex("y")])
				prize = Prize{x, y}
			}
			tmp = append(tmp, Game{buttonA, buttonB, prize})
		}
	}
	return tmp
}

// calculate minium effort to complete the given game
func play(game Game) int {
	for i := 1; i <= 100; i++ {
		for j := 1; j <= 100; j++ {
			if (i*game.buttonA.x+j*game.buttonB.x) == game.prize.x && (i*game.buttonA.y+j*game.buttonB.y) == game.prize.y {
				return i*3 + j
			}
		}
	}
	return 0
}

func play2(game Game) int {
	added := 10000000000000
	game.prize.x += added
	game.prize.y += added

	det := (game.buttonA.x * game.buttonB.y) - (game.buttonB.x * game.buttonA.y)
	if det == 0 {
		return 0 // No solution or infinite solutions
	}

	a := (game.prize.x*game.buttonB.y - game.prize.y*game.buttonB.x) / det
	b := (game.prize.y*game.buttonA.x - game.prize.x*game.buttonA.y) / det

	// Ensure a and b are valid (positive integers)
	if b < 0 || a < 0 ||
		(game.buttonA.x*a+game.buttonB.x*b != game.prize.x) ||
		(game.buttonA.y*a+game.buttonB.y*b != game.prize.y) {
		return 0
	}

	return a*3 + b
}

func part1(input []Game) {
	res := 0
	for _, game := range input {
		res += play(game)
	}
	fmt.Println(res)
}

func part2(input []Game) {
	res := 0
	for _, game := range input {
		res += play2(game)
	}
	fmt.Println(res)
}

func main() {
	input := parse("./input.txt")
	// part1(input)
	part2(input)
}
