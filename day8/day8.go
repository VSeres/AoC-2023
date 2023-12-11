package day8

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	left  string
	right string
}

func Solve(silent bool) {
	file, err := os.Open("day8/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	directions := scanner.Text()

	scanner.Scan()

	nodes := make(map[string]node, 1024)
	startNodes := make([]string, 0, 255)

	for scanner.Scan() {
		line := scanner.Text()
		newNode := node{}
		var key string
		fmt.Sscanf(line, "%3s = (%3s,%3s)", &key, &newNode.left, &newNode.right)
		if key[2] == 'A' {
			startNodes = append(startNodes, key)
		}
		nodes[key] = newNode
	}

	steps := partOne(directions, nodes)
	fmt.Println(steps)

	steps = partTwo(startNodes, nodes, directions)

	fmt.Println(steps)
}

func partOne(directions string, nodes map[string]node) int {
	curNode := nodes["AAA"]
	var (
		steps  = 0
		dirLen = len(directions)
	)
	for {
		dir := directions[steps%dirLen]
		steps++

		var nextNode string
		if dir == 'R' {
			nextNode = curNode.right
		} else {
			nextNode = curNode.left
		}
		if nextNode == "ZZZ" {
			break
		}
		curNode = nodes[nextNode]
	}
	return steps
}

func partTwo(startNodes []string, nodes map[string]node, directions string) int {
	dirLen := len(directions)
	steps := 0
	cycle := make([]int, 0, len(startNodes))
	for _, name := range startNodes {
		n := nodes[name]
		var nextNode string
		zMap := make(map[string]bool, 32)
		steps = 0
		for {
			dir := directions[steps%dirLen]
			steps++
			if dir == 'R' {
				nextNode = n.right
			} else {
				nextNode = n.left
			}
			n = nodes[nextNode]
			if zMap[nextNode] {
				cycle = append(cycle, steps/2)
				break
			} else if nextNode[2] == 'Z' {
				zMap[nextNode] = true
			}
		}

	}
	result := cycle[0]
	for i := 1; i < len(cycle); i++ {
		result = lcm(result, cycle[i])
	}
	return result
}

func lcm(a int, b int) int {
	return (a * b) / gcd(a, b)
}

func gcd(a int, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
