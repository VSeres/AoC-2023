package day1

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

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
		sumPartOne += getNumber(line)
		sumPartTwo += getNumberPEG(line)
	}

	if !silent {
		fmt.Printf("Part one solution: %d\nPart two solution: %d \n", sumPartOne, sumPartTwo)
	}
}

func getNumber(line string) int {
	firstNum := 0
	lastNum := 0
	for _, char := range line {
		if !unicode.IsDigit(char) {
			continue
		}
		num := int(char - '0')
		if firstNum == 0 {
			firstNum += num
		}
		lastNum = num
	}

	return firstNum*10 + lastNum
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
}

func getNumberPEG(line string) int {
	var buff string

	var (
		first int
		last  int
	)

	for _, char := range line {
		var num int

		if unicode.IsDigit(char) {
			buff = ""
			num = int(char - '0')
		} else {
			buff += string(char)
			lookup, ok := numbers[buff]
			// because of gibersih charaters
			if len(buff) > 3 && !ok {
				for i := 1; i <= len(buff)-3; i++ {
					lookup, ok = numbers[buff[i:]]
					if ok {
						break
					}
				}
			}
			if ok {
				num = lookup
				// Keep the last letter, because another one can start with it
				buff = buff[len(buff)-1:]
			}
		}

		if num == 0 {
			continue
		} else if first == 0 {
			first = num
		}
		last = num

	}

	return first*10 + last
}
