package day11

import (
	"bufio"
	"fmt"
	"os"
)

type galaxy struct {
	x int
	y int
}

func Solve(silent bool) {
	file, err := os.Open("day11/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	galaxyMap := make([][]byte, 0)
	galaxiesInColoumn := make(map[int]int, 0)
	galaxiesInRow := make(map[int]int, 0)

	galaxies := make([]galaxy, 0)

	y := 0
	for scanner.Scan() {
		line := scanner.Bytes()
		mapLine := make([]byte, len(line))
		copy(mapLine, line)
		galaxyMap = append(galaxyMap, mapLine)
		for x, place := range mapLine {
			if place != '.' {
				galaxiesInColoumn[x] += 1
				galaxiesInRow[y] += 1
				galaxies = append(galaxies, galaxy{x: x, y: y})
			}
		}
		y++
	}

	xOffset := make([]int, len(galaxyMap[0]))
	yOffset := make([]int, len(galaxyMap))
	// calculate x offset
	for x := 0; x < len(xOffset); x++ {
		if galaxiesInColoumn[x] == 0 {
			for i := x; i < len(xOffset); i++ {
				xOffset[i] += 1
			}
		}
	}
	// calculate y offset
	for y := 0; y < len(yOffset); y++ {
		if galaxiesInRow[y] == 0 {
			for i := y; i < len(yOffset); i++ {
				yOffset[i] += 1
			}
		}
	}

	for i, g := range galaxies {
		galaxies[i].x += xOffset[g.x]
		galaxies[i].y += yOffset[g.y]
	}
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += manhattanDistance(galaxies[i], galaxies[j])
		}
	}
	fmt.Println(sum)
}

func manhattanDistance(g1, g2 galaxy) int {
	x := g1.x - g2.x
	y := g1.y - g2.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
