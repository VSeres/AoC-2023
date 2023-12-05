package main

import (
	"AoC-2023/day1"
	"AoC-2023/day2"
	"AoC-2023/day3"
	"AoC-2023/day4"
	"AoC-2023/day5"
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
		day3.Solve()
	case 4:
		day4.Solve()
	case 5:
		day5.Solve()
	default:
		fmt.Printf("day %d not implemented\n", day)
	}

}
