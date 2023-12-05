package day3

import (
	"bufio"
	"fmt"
	"os"
)

func Solve(silent bool) {
	file, err := os.Open("day3/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	engine := make([][]byte, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		engine = append(engine, []byte(line))
	}
	partOne(engine, silent)
}

func specialChar(char byte) bool {
	return (char < '0' || char > '9') && char != '.'
}

type point struct {
	x int
	y int
}

type symbol struct { // for part two
	x    int
	y    int
	char byte
}

func partOne(engine [][]byte, silent bool) {
	gears := make(map[point][]int, 0)
	result := 0
	for y, line := range engine {
		partNumber := 0
		var special *symbol
		for x, char := range line {
			if char >= '0' && char <= '9' {
				partNumber = partNumber*10 + int(char-'0')
				// right
				if x < len(line)-1 && specialChar(line[x+1]) {
					special = &symbol{x + 1, y, line[x+1]}
				}
				// left
				if x > 0 && specialChar(line[x-1]) {
					special = &symbol{x - 1, y, line[x-1]}
				}
				// up
				if y > 0 && specialChar(engine[y-1][x]) {
					special = &symbol{x, y - 1, engine[y-1][x]}
				}
				// down
				if y < len(engine)-1 && specialChar(engine[y+1][x]) {
					special = &symbol{x, y + 1, engine[y+1][x]}
				}
				// top right
				if x > 0 && y > 0 && specialChar(engine[y-1][x-1]) {
					special = &symbol{x - 1, y - 1, engine[y-1][x-1]}
				}
				// top left
				if x < len(line)-1 && y > 0 && specialChar(engine[y-1][x+1]) {
					special = &symbol{x + 1, y - 1, engine[y-1][x+1]}
				}
				// bottom right
				if x > 0 && y < len(engine)-1 && specialChar(engine[y+1][x-1]) {
					special = &symbol{x - 1, y + 1, engine[y+1][x-1]}
				}
				// bottom left
				if x < len(line)-1 && y < len(engine)-1 && specialChar(engine[y+1][x+1]) {
					special = &symbol{x + 1, y + 1, engine[y+1][x+1]}
				}

			} else if special != nil {
				// fmt.Println(n)
				result += partNumber
				if special.char == '*' {
					p := point{x: special.x, y: special.y}
					_, ok := gears[p]
					if ok {
						gears[p] = append(gears[p], partNumber)
					} else {
						gears[p] = []int{partNumber}
					}
				}
				special = nil
				partNumber = 0
			} else {
				// if n > 0 {
				// 	fmt.Println(n)
				// }
				partNumber = 0
				special = nil
			}
		}
		if special != nil {
			result += partNumber
			if special.char == '*' {
				p := point{x: special.x, y: special.y}
				_, ok := gears[p]
				if ok {
					gears[p] = append(gears[p], partNumber)
				} else {
					gears[p] = []int{partNumber}
				}
			}
		}
	}
	// part two
	sum := 0
	for _, partNumbers := range gears {
		if len(partNumbers) != 2 {
			continue
		}
		sum += partNumbers[0] * partNumbers[1]
	}
	if !silent {
		fmt.Println(result)
		fmt.Println(sum)
	}

}
