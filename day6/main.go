package day6

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func Solve(silent bool) {
	file, err := os.Open("day6/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// read time
	scanner.Scan()
	timeLine := scanner.Text()
	cutIndex := strings.Index(timeLine, ":") + 1
	timeLine = strings.Trim(timeLine[cutIndex:], " ")
	time := make([]int, 0)
	num := 0

	// part two
	var (
		bigTime     int
		bigDistance int
	)
	//---
	for _, v := range timeLine {
		if v != ' ' {
			num = num*10 + int(v-'0')
			bigTime = bigTime*10 + int(v-'0')
		} else if num > 0 {
			time = append(time, num)
			num = 0
		}
	}
	if num > 0 {
		time = append(time, num)
	}
	// read distance
	scanner.Scan()
	distanceLine := scanner.Text()
	cutIndex = strings.Index(distanceLine, ":") + 1
	distanceLine = strings.Trim(distanceLine[cutIndex:], " ")
	distance := make([]int, 0)
	num = 0
	for _, v := range distanceLine {
		if v != ' ' {
			num = num*10 + int(v-'0')
			bigDistance = bigDistance*10 + int(v-'0')
		} else if num > 0 {
			distance = append(distance, num)
			num = 0
		}
	}
	if num > 0 {
		distance = append(distance, num)
	}
	margin := 1
	for i, t := range time {
		d := distance[i]
		margin *= calculateMargin(t, d)
	}

	if !silent {
		fmt.Printf("Margin of error: %d\nFix margin: %d\n", margin, calculateMargin(bigTime, bigDistance))
	}

}

func calculateMargin(t int, d int) int {
	// +x2 - nx + c = 0
	root := math.Sqrt(float64(t*t - 4*d))
	x1 := (float64(t) - root) * 0.5
	x2 := (float64(t) + root) * 0.5
	start := int(x1) + 1
	end := int(math.Ceil(x2)) - 1

	return end - start + 1
}
