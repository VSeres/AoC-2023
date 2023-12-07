package day7

import (
	"bufio"
	"fmt"
	"os"
)

type match struct {
	hand string
	bid  int
	rank byte
}

const cardCount = 5

func Solve(silent bool) {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	matches := make([]match, 0, 1024)

	for scanner.Scan() {
		line := scanner.Text()
		var opt match
		fmt.Sscan(line, &opt.hand, &opt.bid)

		repeating := [2]byte{}
		var firstMatch byte
		for i := 0; i < cardCount-1; i++ {
			for j := i + 1; j < cardCount; j++ {
				if opt.hand[i] == opt.hand[j] {
					if firstMatch == 0 {
						firstMatch = opt.hand[i]
					}
					if opt.hand[i] == firstMatch {
						repeating[0] += 1
					} else {
						repeating[1] += 1
					}

					break
				}
			}
		}

		if repeating[0] < repeating[1] {
			repeating[0], repeating[1] = repeating[1], repeating[0]
		}

		repeating[0] += 1
		repeating[1] += 1

		if repeating[0] == 5 {
			opt.rank = 1
		} else if repeating[0] == 4 {
			opt.rank = 2
		} else if repeating[0] == 3 && repeating[1] == 2 {
			opt.rank = 3
		} else if repeating[0] == 3 {
			opt.rank = 4
		} else if repeating[0] == 2 && repeating[1] == 2 {
			opt.rank = 5
		} else if repeating[0] == 2 {
			opt.rank = 6
		} else {
			opt.rank = 7
		}

		matches = append(matches, opt)
	}
	sort(matches)
	winnings := 0

	for rank, opt := range matches {
		winnings += (rank + 1) * opt.bid
	}
	if !silent {
		fmt.Println(winnings)
	}
}

func sort(matches []match) {
	n := len(matches)
	newn := 0
	swap := true
	for swap {
		swap = false
		for i := 1; i <= n-1; i++ {
			if compare(matches[i-1], matches[i]) {
				matches[i-1], matches[i] = matches[i], matches[i-1]
				newn = i
				swap = true
			}
		}
		n = newn
	}
}

func compare(a match, b match) bool {
	if a.rank < b.rank {
		return true
	} else if a.rank == b.rank {
		for i := 0; i < 5; i++ {
			aValue := getValue(a.hand[i])
			bValue := getValue(b.hand[i])
			if aValue > bValue {
				return true
			} else if bValue > aValue {
				return false
			}
		}
		panic("same")
	}
	return false
}

func getValue(a byte) byte {
	switch a {
	case 'T':
		return 58
	case 'J':
		return 59
	case 'Q':
		return 60
	case 'K':
		return 61
	case 'A':
		return 62
	default:
		return a
	}
}
