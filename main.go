package main

import (
	"AoC-2023/day1"
	"AoC-2023/day2"
	"AoC-2023/day3"
	"flag"
	"fmt"
)

var day int

func init() {
	flag.IntVar(&day, "day", 1, "Which day should be run.")
}

func main() {
	flag.Parse()

	switch day {
	case 1:
		day1.Solve()
	case 2:
		day2.Solve()
	case 3:
		day3.Solver()
	default:
		fmt.Printf("day %d not implemented\n", day)
	}

}
