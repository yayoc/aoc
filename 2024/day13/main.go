package main

type Button {
    x int
    y int
}

type Prize {
    x int
    y int
}

func parse(filePath string) [](Button, Button, Prize) {
    file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
    var buttonA: Button
    var buttonB: Button
    var prize: Prize
    tmp := make([](Button, Button, Prize), 0)
	for scanner.Scan() {
		line := scanner.Text()
        fields := strings.Fields(line)
        if len(fields) == 0 {
            tmp = make([](Button, Button, Prize), 0)
        }

	}
}

func main() {
}
