package day1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func Solve() {
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
	count := 1
	for scanner.Scan() {
		line := scanner.Text()
		sumPartOne += getNumber(line)
		num := getNumberPEG(line)
		sumPartTwo += num
		count++
	}
	fmt.Printf("Part one solution: %d\nPart two solution: %d \n", sumPartOne, sumPartTwo)
}

func getNumber(line string) int {
	numberString := ""
	var lastChar rune
	for _, char := range line {
		if !unicode.IsDigit(char) {
			continue
		}

		if numberString == "" {
			numberString += string(char)
		}
		lastChar = char
	}

	numberString += string(lastChar)

	num, err := strconv.Atoi(numberString)
	if err != nil {
		panic(err)
	}
	return num
}

var numbers map[string]string = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func getNumberPEG(line string) int {
	var buff string

	var (
		first string
		last  string
	)

	for _, char := range line {
		var num string

		if unicode.IsDigit(char) {
			buff = ""
			num = string(char)
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

		if num == "" {
			continue
		} else if first == "" {
			first = num
		}
		last = num

	}

	result, err := strconv.Atoi(first + last)
	if err != nil {
		panic(err)
	}
	return result
}
