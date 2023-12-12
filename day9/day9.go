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
	futureSum := 0
	pastSum := 0
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
			nextSeriesLen := len(c) - 1
			arr := make([]int, nextSeriesLen)
			zeroCount := calcChanges(c, arr)
			finished = zeroCount == nextSeriesLen
			if nextSeriesLen == 0 {
				break
			}
			changes = append(changes, arr)
			i++
		}

		// calculate last element

		futureSum += extrapolateFutuere(changes)
		pastSum += extrapolatePast(changes)
		prettyPrint(changes)
	}
	fmt.Println(futureSum, pastSum)
}

func extrapolateFutuere(changes [][]int) int {
	prevChange := 0
	for j := len(changes) - 1; j >= 0; j-- {
		prevChange += changes[j][len(changes[j])-1]
	}

	return prevChange
}

func extrapolatePast(changes [][]int) int {
	prevChange := 0
	for j := len(changes) - 1; j >= 0; j-- {
		prevChange = changes[j][0] - prevChange
	}

	return prevChange
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

const numPadding = 2

func prettyPrint(nums [][]int) {
	builder := new(strings.Builder)
	maxLen := len(nums[0])
	for _, line := range nums {
		padding := maxLen - len(line)
		if padding > 0 {
			builder.WriteString(strings.Repeat(" ", padding*numPadding))
		}
		for _, num := range line {
			fmt.Fprintf(builder, "%2d  ", num)
		}
		builder.WriteByte('\n')
	}
	fmt.Println(builder.String())
}

func readNumbers(scanner *bufio.Scanner) []int {
	n := 0
	change := make([]int, 0)
	negative := false
	for _, v := range scanner.Bytes() {
		if v == ' ' {
			if negative {
				n = -n
			}
			change = append(change, n)
			negative = false
			n = 0
		} else if v == '-' {
			negative = true
		} else {
			n = n*10 + int(v-'0')
		}
	}
	if negative {
		n = -n
	}
	change = append(change, n)
	return change
}
