package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var maxCount = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}

func Solve() {
	file, err := os.Open("day2/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	gameId := 1

	result := 0   // part one result
	powerSum := 0 // part two result

	for scanner.Scan() {
		line := scanner.Text()
		i := strings.IndexRune(line, ':')
		line = line[i+2:]
		game := strings.Split(line, " ")

		cubesNeeded := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}

		vailid := true
		for i := 0; i < len(game); i += 2 {
			pullCount := game[i]
			// remove trailing ',' and ';' from colors
			pullColor := strings.Trim(game[i+1], ",;")

			count, err := strconv.Atoi(pullCount)
			if err != nil {
				panic(err)
			}
			if count > maxCount[pullColor] {
				// fmt.Printf("%s maximum is %d but pulled %d in game %d\n", pullColor, maxCount[pullColor], count, gameId)
				vailid = false
			}

			if cubesNeeded[pullColor] < count {
				cubesNeeded[pullColor] = count
			}
		}
		//upadte part one result
		if vailid {
			result += gameId
		}

		//upadte part two result
		power := cubesNeeded["red"] * cubesNeeded["green"] * cubesNeeded["blue"]
		// fmt.Printf("Power %d, gameId: %d\n", power, gameId)
		powerSum += power

		gameId++
	}
	fmt.Println("Part one result: ", result)
	fmt.Println("Part two result: ", powerSum)
}
