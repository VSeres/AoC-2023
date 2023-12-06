package day1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var digit = regexp.MustCompile(`\d`)

func Solve(silent bool) {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var (
		sumPartOne int
		sumPartTwo int
	)

	for scanner.Scan() {
		line := scanner.Text()
		sumPartOne += partOne(line)

		sumPartTwo += partTwo(line)
	}

	if !silent {
		fmt.Printf("Part one solution: %d\nPart two solution: %d \n", sumPartOne, sumPartTwo)
	}
}

func partTwo(line string) int {
	lastIndex := -999
	lastValue := 0
	firstIndex := 999
	firstValue := 0
	for k, v := range numbers {
		firstI := strings.Index(line, k)
		lastI := strings.LastIndex(line, k)
		if lastI > lastIndex {
			lastIndex = lastI
			lastValue = v
		}
		if firstI >= 0 && firstI < firstIndex {
			firstIndex = firstI
			firstValue = v
		}
	}
	res := firstValue*10 + lastValue
	return res
}

func partOne(line string) int {
	match := digit.FindAllString(line, -1)
	first := int(match[0][0] - '0')
	last := int(match[len(match)-1][0] - '0')
	return first*10 + last
}

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
	"1":     1,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
}
