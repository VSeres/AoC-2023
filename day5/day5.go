package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	destination int
	source      int
	span        int
}

func (r rule) apply(factor int) int {
	upperBound := r.source + r.span - 1
	if r.source <= factor && factor <= upperBound {
		return factor - r.source + r.destination
	}
	return factor
}

func (r rule) revers(dest int) int {
	if r.destination+r.span <= dest || r.destination > dest {
		return dest
	}
	return dest - r.destination + r.source
}

func Solve(silent bool) {
	file, err := os.Open("day5/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// read seeds
	seeds := make([]int, 0)
	scanner.Scan()

	seedLine := scanner.Text()
	cutIndex := strings.Index(seedLine, ":") + 2
	seedLine = seedLine[cutIndex:]

	for _, v := range strings.Split(seedLine, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, num)
	}

	scanner.Scan() // empty line

	rules := make([][]rule, 7)

	seedsCopy := make([]int, len(seeds))
	copy(seedsCopy, seeds)

	partOne := lowestLocation(scanner, seedsCopy, rules)

	partTwo := lowestLocasonBackwards(rules, seeds)

	if !silent {
		fmt.Printf("Lowest location in:\n\tpart one: %d\n\tpart two: %d\n", partOne, partTwo)
	}
}

func lowestLocasonBackwards(rules [][]rule, sed []int) int {

	for i, j := 0, len(rules)-1; i < j; i, j = i+1, j-1 {
		rules[i], rules[j] = rules[j], rules[i]
	}

	location := 0
	var num int
	for !inRange(sed, num) {
		location++
		num = location

		for _, ruleList := range rules {
			for _, rule := range ruleList {

				newNum := rule.revers(num)
				if newNum != num {
					num = newNum
					break
				}
			}
		}
	}
	return location
}

func inRange(seed []int, num int) bool {
	for i := 0; i < len(seed); i += 2 {
		if num >= seed[i] && seed[i]+seed[i+1] > num {
			return true
		}
	}
	return false
}

func lowestLocation(scanner *bufio.Scanner, seeds []int, rules [][]rule) int {
	ruleList := make([]rule, 0)
	state := 0
	for state < 7 {
		scanner.Scan()

		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			var newRule rule
			fmt.Sscan(line, &newRule.destination, &newRule.source, &newRule.span)
			ruleList = append(ruleList, newRule)
		}

		ruleListCopy := make([]rule, len(ruleList))
		copy(ruleListCopy, ruleList)

		rules[state] = ruleListCopy

		for sedIndex, sed := range seeds {
			for _, rule := range ruleList {
				num := rule.apply(sed)
				seeds[sedIndex] = num
				if num != sed {
					break
				}
			}
		}
		state++

		ruleList = ruleList[:0]
	}
	minIndex := 0
	for index, s := range seeds {
		if seeds[minIndex] > s {
			minIndex = index
		}
	}
	return seeds[minIndex]
}
