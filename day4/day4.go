package day4

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Solve(silent bool) {
	file, err := os.Open("day4/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := 0

	cardCount := make([]int, 255)
	currentCard := 0
	for scanner.Scan() {
		line := scanner.Text()
		i := strings.Index(line, ":") + 1
		line = line[i:]

		numbers := strings.Split(line, "|")
		numbers[0] = strings.Trim(numbers[0], " ")
		numbers[1] = strings.Trim(numbers[1], " ")

		guess := make(map[string]bool, 0)

		for _, v := range strings.Split(numbers[0], " ") {
			if v == "" {
				continue
			}
			guess[v] = true
		}

		winningNums := make(map[string]bool, 0)

		for _, v := range strings.Split(numbers[1], " ") {
			if v == "" {
				continue
			}
			winningNums[v] = true
		}

		matchCount := 0

		for k := range guess {
			if winningNums[k] {
				matchCount++
			}
		}

		if cardCount[currentCard] == 0 {
			cardCount[currentCard] = 1
		}

		for i := currentCard + 1; i <= currentCard+matchCount; i++ {
			if cardCount[i] == 0 {
				cardCount[i] = 1
			}
			cardCount[i] += cardCount[currentCard]
		}
		result += int(math.Pow(2, float64(matchCount-1)))
		currentCard++
	}
	sum := 0
	for _, v := range cardCount {
		sum += v
	}
	if !silent {
		fmt.Printf("Part one solustion: %d\nTotal number of cards: %d\n", result, sum)
	}
}
