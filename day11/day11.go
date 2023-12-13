package day11

import (
	"bufio"
	"fmt"
	"os"
)

type galaxy struct {
	x          int
	y          int
	alternateX int
	alternateY int
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
		// part two
		galaxies[i].alternateX = g.x + xOffset[g.x]*(1000000-1) // ther are 1000000-1 new lines
		galaxies[i].alternateY = g.y + yOffset[g.y]*(1000000-1)
	}

	youngSum := 0
	oldSum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			g1 := galaxies[i]
			g2 := galaxies[j]
			youngSum += manhattanDistance(g1.x, g2.x, g1.y, g2.y)
			oldSum += manhattanDistance(g1.alternateX, g2.alternateX, g1.alternateY, g2.alternateY)
		}
	}
	fmt.Println(youngSum, oldSum)
}

func manhattanDistance(x1, x2, y1, y2 int) int {
	x := x1 - x2
	y := y1 - y2
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
