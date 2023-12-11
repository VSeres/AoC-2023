package day9

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Solve(silent bool) {
	file, err := os.Open("day9/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	valueSum := 0
	for scanner.Scan() {
		changes := make([][]int, 0)
		changeSeries := readNumbers(scanner)
		changes = append(changes, changeSeries)
		finished := false
		i := 0
		for !finished {
			c := changes[i]
			// if len(c) == 1 {
			// 	changes = append(changes, make([]int, 1))
			// 	break
			// }
			arr := make([]int, len(c)-1)
			zeroCount := calcChanges(c, arr)
			changes = append(changes, arr)
			finished = zeroCount == len(changes[i+1])
			i++
		}
		// calculate last element
		prevChange := 0
		for j := len(changes) - 2; j >= 0; j-- {
			prevChange += changes[j][len(changes[j])-1]
		}
		valueSum += prevChange
		prettyPrint(changes)
	}
	fmt.Println(valueSum)
}

func calcChanges(c []int, arr []int) int {
	prevChange := 0
	for j := 1; j < len(c); j++ {
		change := c[j] - c[j-1]
		if change == 0 {
			prevChange++
		}
		arr[j-1] = change
	}
	return prevChange
}

const numPadding = 4

func prettyPrint(nums [][]int) {
	builder := new(strings.Builder)
	maxLen := len(nums[0])
	for _, line := range nums {
		padding := maxLen - len(line)
		if padding > 0 {
			builder.WriteString(strings.Repeat(" ", padding*numPadding))
		}
		for _, num := range line {
			fmt.Fprintf(builder, "%4d  ", num)
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func readNumbers(scanner *bufio.Scanner) []int {
	n := 0
	change := make([]int, 0)
	for _, v := range scanner.Bytes() {
		if v == ' ' {
			change = append(change, n)
			n = 0
		} else {
			n = n*10 + int(v-'0')
		}
	}
	change = append(change, n)
	return change
}
