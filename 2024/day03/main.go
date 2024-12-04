package main

func isNum(s string) bool {
    _, err := strconv.ParseFloat(s, 64)
    return err == nil
}

func parse(filePath string) [](int,int) {
    b, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
    str := string(b)
    var nums [](int, int)

    match, _ := regexp.MatchString("mul\(\d{1,3},\d{1,3}\)", str)
    fmt.Println(match)
}

func part1(nums [](int,int)) {
    sum := 0
    for x, y := range nums {
        sum += x * y
    }
    fmt.Println(sum)
}

func main() {
    nums := parse("./input.txt")
    part1(nums)
}
