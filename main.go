package main

import (
	"AoC-2023/day1"
	"AoC-2023/day2"
	"AoC-2023/day3"
	"AoC-2023/day4"
	"AoC-2023/day5"
	"AoC-2023/day6"
	"flag"
	"fmt"
	"time"
)

var day int

func init() {
	flag.IntVar(&day, "day", 1, "Which day should be run.")
}

func main() {
	flag.Parse()
	start := time.Now()
	switch day {
	case 1:
		day1.Solve(false)
	case 2:
		day2.Solve(false)
	case 3:
		day3.Solve(false)
	case 4:
		day4.Solve(false)
	case 5:
		day5.Solve(false)
	case 6:
		day6.Solve(false)
	default:
		fmt.Printf("day %d not implemented\n", day)
	}

	defer fmt.Printf("Executon took %s\n", time.Since(start))
}
