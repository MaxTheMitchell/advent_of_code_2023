package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input := readInput("./input.txt")

	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		sum += getNumb(line)
	}

	fmt.Println(sum)
}

func getNumb(line string) int {
	startNumb := -1
	lastNumb := 0
	for len(line) > 0 {
		ok, num := nextNumb(line)
		if ok {
			lastNumb = num
			if startNumb == -1 {
				startNumb = num
			}
		}
		line = line[1:]
	}
	if startNumb == -1 {
		startNumb = 0
	}

	return startNumb*10 + lastNumb
}

func nextNumb(line string) (bool, int) {
	numbMap := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	num, err := strconv.ParseInt(string(rune(line[0])), 10, 64)
	if err == nil {
		return true, int(num)
	}

	for numStr, weight := range numbMap {
		if strings.HasPrefix(line, numStr) {
			return true, weight
		}
	}

	return false, 0
}

func readInput(fileName string) string {
	bytes, _ := os.ReadFile(fileName)
	return string(bytes)
}
