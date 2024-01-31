package main

import (
	"AoC-2023/day1"
	"AoC-2023/day10"
	"AoC-2023/day11"
	"AoC-2023/day12"
	"AoC-2023/day2"
	"AoC-2023/day3"
	"AoC-2023/day4"
	"AoC-2023/day5"
	"AoC-2023/day6"
	"AoC-2023/day7"
	"AoC-2023/day8"
	"AoC-2023/day9"
	"flag"
	"fmt"
	"time"
)

var day int

func init() {
	flag.IntVar(&day, "day", 1, "Which day should be run.")
}

func main() {
	// day10.Solve(false)
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
	case 7:
		day7.Solve(false)
	case 8:
		day8.Solve(false)
	case 9:
		day9.Solve(false)
	case 10:
		day10.Solve(false)
	case 11:
		day11.Solve(false)
	case 12:
		day12.Solve()
	default:
		fmt.Printf("day %d not implemented\n", day)
	}

	defer fmt.Printf("Executon took %s\n", time.Since(start))
}
