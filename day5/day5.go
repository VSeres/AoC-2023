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

func Solve() {
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
	i := strings.Index(seedLine, ":") + 2
	seedLine = seedLine[i:]

	for _, v := range strings.Split(seedLine, " ") {
		num, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, num)
	}

	scanner.Scan() // empty line

	partOne := lowestLocation(scanner, seeds)

	// partTwoSeeds := make([]int, 0)
	// for i := 0; i < len(seeds); i += 2 {
	// 	for seed := seeds[i]; seed < seeds[i]+seeds[i+1]; seed++ {
	// 		partTwoSeeds = append(partTwoSeeds, seed)
	// 	}
	// }
	// parTwo := lowestLocation(scanner, partTwoSeeds)

	fmt.Printf("Lowest location in:\n\tpart one: %d\n\tpart two: %d\n", partOne, -1)
}

func lowestLocation(scanner *bufio.Scanner, seeds []int) int {
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
