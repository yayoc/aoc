package main

import (
    "os"
    "log"
    "fmt"
    "strings"
    "strconv"
    "math"
)

func parse(filePath string) []int {
    b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	str := string(b)
    fields := strings.Fields(str)
    res := make([]int, 0)
    for _, f := range fields {
        i, _ := strconv.Atoi(f)
        res = append(res, i)
    }
    return res
}

// to insert nums into index i of slice
func insert(slice []int, i int, nums []int) []int {
    l := slice[:i]
    r := slice[i:]

    return append(append(l, nums...), r...)
}

// Check the number of digits without string conversion
func numDigits(num int) int {
    if num == 0 {
        return 1
    }
    return int(math.Floor(math.Log10(float64(num)))) + 1
}

func divide(num int) (int, int) {
    digits := int(math.Log10(float64(num))) + 1
    half := digits / 2
    divisor := int(math.Pow10(half))

    left := num / divisor
    right := num % divisor
    return left, right
}


// tried to solve part2 updating slice in-place approach, but didn't work
func part1(nums []int, count int) {
    n := 0
    for n < count {
        total := 0
        // calc total length
        for _, num := range nums {
            if numDigits(num) % 2 == 0 {
                total += 2
            } else {
                total++
            }
        }

        // extend nums
        if cap(nums) < total {
            newSlice := make([]int, len(nums), total)
            copy(newSlice, nums)
            nums = newSlice
        }

        oldLen := len(nums)
        nums = nums[:total]

        writePos := total - 1

        // update nums in-place
        for i := oldLen - 1; i >= 0; i-- {
            num := nums[i]
            if num == 0 {
                nums[writePos] = 1
                writePos--
            } else {
                if numDigits(num) % 2 == 0 {
                    l, r := divide(num)
                    nums[writePos] = r
                    writePos--
                    nums[writePos] = l
                    writePos--
                } else {
                    nums[writePos] = num * 2024
                    writePos--
                }
            }
        }
        n++
    }
    fmt.Println(len(nums))
}

func part2(nums []int) {
    memo := make(map[int]int, 0)
    // initialize
    for _, num := range nums {
        memo[num] = 1
    }
    n := 0
    for n < 75 {
        old := make(map[int]int, 0)
        for k, v := range memo {
            old[k] = v
        }
        for num, count := range old {
            if num == 0 {
                memo[1] += count
                memo[0] -= count
            } else {
                if numDigits(num) % 2 == 0 {
                    l, r := divide(num)
                    memo[num] -= count
                    memo[l] += count
                    memo[r] += count
                } else {
                    memo[num*2024] += count
                    memo[num] -= count
                }
            }
        }
        n++
    }

    res := 0
    for _, v := range memo {
        res += v
    }
    fmt.Println(res)
}


func main() {
    nums := parse("./input.txt")
    // part1(nums, 25)
    part2(nums)
}
