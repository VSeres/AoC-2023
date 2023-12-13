package day10

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type node struct {
	key           int
	distance      int
	adjacencyList []*node
}

const (
	lineLen   = 140
	lineCount = 140
)

type direction byte

const (
	none direction = iota
	north
	east
	south
	west
)

var pipeMap = make([][]byte, lineCount)
var visited = make(map[int]int, lineCount*lineLen)

func Solve(silent bool) {
	file, err := os.Open("day10/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var (
		sX int
		sY int
		y  int
	)

	for scanner.Scan() {
		line := scanner.Bytes()
		pipeMap[y] = make([]byte, lineLen)
		copy(pipeMap[y], line)
		for x, v := range line {
			if v == 'S' {
				sX = x
				sY = y
			}
		}
		y++
	}

	travers(0, sX, sY, none)
	maxD := 0
	for _, v := range visited {
		if v > maxD {
			maxD = v
		}
	}
	// var b strings.Builder
	// b.Grow(lineCount*lineLen + lineCount)

	expand()
	flood()
	collaps()
	// printMap(&b, expandMap)
	// b.WriteByte('\n')
	// printMap(&b, collapesMap)
	// fmt.Println(b.String())
	fmt.Printf("max distance: %d\nenclosed spaces: %d\n", maxD, countSpaces())
}

func countSpaces() int {
	empty := 0
	for _, line := range collapesMap {
		for _, place := range line {
			if place == 0 {
				empty += 1
			}
		}
	}
	return empty
}

func printMap(b *strings.Builder, m [][]byte) {
	for _, line := range m {
		for _, pipe := range line {
			if pipe == 0 {
				b.WriteByte('?')
			} else {
				b.WriteByte(pipe)
			}
		}
		b.WriteByte('\n')
	}
}

func keyFunc(a int, b int) int {
	top := (a + b) * (a + b + 1)
	return top/2 + a
}

func travers(d int, x int, y int, dir direction) {
	key := keyFunc(x, y)
	value, ok := visited[key]
	if !ok {
		visited[key] = d
	} else if value > d {
		visited[key] = d
	} else {
		return
	}

	switch dir {
	case north:
		if y == 0 {
			break
		}
		newDir := checkNorth(x, y)
		if newDir != none {
			travers(d+1, x, y-1, newDir)
		}
	case east:
		if x == lineLen-1 {
			break
		}
		newDir := checkEast(x, y)
		if newDir != none {
			travers(d+1, x+1, y, newDir)
		}
	case south:
		if y == lineCount-1 {
			break
		}
		newDir := checkSouth(x, y)
		if newDir != none {
			travers(d+1, x, y+1, newDir)
		}
	case west:
		if x == 0 {
			break
		}
		newDir := checkWest(x, y)
		if newDir != none {
			travers(d+1, x-1, y, newDir)
		}
	case none:
		var newDir direction
		if y > 0 {
			newDir = checkNorth(x, y)
			if newDir != none {
				travers(d+1, x, y-1, newDir)
			}
		}
		if x < lineLen-1 {
			newDir = checkEast(x, y)
			if newDir != none {
				travers(d+1, x+1, y, newDir)
			}
		}
		if y < lineCount-1 {
			newDir = checkSouth(x, y)
			if newDir != none {
				travers(d+1, x, y+1, newDir)
			}
		}
		if x > 0 {
			newDir = checkWest(x, y)
			if newDir != none {
				travers(d+1, x-1, y, newDir)
			}
		}
	}

}

func checkNorth(x, y int) direction {
	switch pipeMap[y-1][x] {
	case '|':
		return north
	case '7':
		return west
	case 'F':
		return east
	default:
		return none
	}
}

func checkEast(x, y int) direction {
	switch pipeMap[y][x+1] {
	case '-':
		return east
	case '7':
		return south
	case 'J':
		return north
	default:
		return none
	}
}

func checkSouth(x, y int) direction {
	switch pipeMap[y+1][x] {
	case '|':
		return south
	case 'L':
		return east
	case 'J':
		return west
	default:
		return none
	}
}

func checkWest(x, y int) direction {
	switch pipeMap[y][x-1] {
	case '-':
		return west
	case 'L':
		return north
	case 'F':
		return south
	default:
		return none
	}
}

var expandMap = make([][]byte, lineCount*3)

func expand() {
	for i := 0; i < len(expandMap); i++ {
		expandMap[i] = make([]byte, lineLen*3)
	}
	for y, line := range pipeMap {
		for x, pipe := range line {
			if _, ok := visited[keyFunc(x, y)]; !ok {
				continue
			}
			eX := x * 3
			eY := y * 3
			switch pipe {
			case '|':
				expandMap[eY][eX+1] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+2][eX+1] = 'X'
			case '-':
				expandMap[eY+1][eX] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+1][eX+2] = 'X'
			case 'L':
				expandMap[eY][eX+1] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+1][eX+2] = 'X'
			case 'J':
				expandMap[eY][eX+1] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+1][eX] = 'X'
			case '7':
				expandMap[eY+1][eX] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+2][eX+1] = 'X'
			case 'F':
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+1][eX+2] = 'X'
				expandMap[eY+2][eX+1] = 'X'
			case 'S':
				expandMap[eY][eX+1] = 'X'
				expandMap[eY+2][eX+1] = 'X'
				expandMap[eY+1][eX+1] = 'X'
				expandMap[eY+1][eX+0] = 'X'
				expandMap[eY+1][eX+2] = 'X'
			}
		}
	}
}

var collapesMap = make([][]byte, 0, lineCount)

func collaps() {
	for y := 1; y < lineCount*3; y += 3 {

		lineMap := make([]byte, 0, lineLen)
		for x := 1; x < lineLen*3; x += 3 {
			lineMap = append(lineMap, expandMap[y][x])
		}
		collapesMap = append(collapesMap, lineMap)
	}

}

type point struct {
	x int
	y int
}

func flood() {
	shouldFlood := make([]point, 0)
	for x, pipe := range expandMap[0] {
		if pipe == 0 {
			shouldFlood = append(shouldFlood, point{x: x, y: 0})
		}
	}
	for x, pipe := range expandMap[lineCount*3-1] {
		if pipe == 0 {
			shouldFlood = append(shouldFlood, point{x: x, y: lineCount*3 - 1})
		}
	}

	p := shouldFlood[len(shouldFlood)-1]
	expandMap[p.y][p.x] = '#'

	for len(shouldFlood) > 0 {
		last := len(shouldFlood) - 1
		pos := shouldFlood[last]
		shouldFlood = shouldFlood[:last]
		if pos.x > 0 && expandMap[pos.y][pos.x-1] == 0 {
			shouldFlood = append(shouldFlood, point{pos.x - 1, pos.y})
			expandMap[pos.y][pos.x-1] = '#'
		}
		if pos.y > 0 && expandMap[pos.y-1][pos.x] == 0 {
			shouldFlood = append(shouldFlood, point{pos.x, pos.y - 1})
			expandMap[pos.y-1][pos.x] = '#'
		}

		if pos.x < lineLen*3-1 && expandMap[pos.y][pos.x+1] == 0 {
			shouldFlood = append(shouldFlood, point{pos.x + 1, pos.y})
			expandMap[pos.y][pos.x+1] = '#'
		}
		if pos.y < lineCount*3-1 && expandMap[pos.y+1][pos.x] == 0 {
			shouldFlood = append(shouldFlood, point{pos.x, pos.y + 1})
			expandMap[pos.y+1][pos.x] = '#'
		}
	}
}

// | is a vertical pipe connecting north and south.
// - is a horizontal pipe connecting east and west.
// L is a 90-degree bend connecting north and east.
// J is a 90-degree bend connecting north and west.
// 7 is a 90-degree bend connecting south and west.
// F is a 90-degree bend connecting south and east.
// . is ground; there is no pipe in this tile.
