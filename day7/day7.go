package day7

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type round struct {
	Hand     string
	Bet      int
	Strength strength
	partTwo  strength
}

type strength int

func Solve(silent bool) {
	var err error
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hands := make([]round, 0)

	for {
		r := round{}
		_, err = fmt.Fscanf(file, "%s %d\n", &r.Hand, &r.Bet)
		if err != nil {
			break
		}
		r.Strength = calcStrength(r.Hand, false)
		r.partTwo = calcStrength(r.Hand, true)

		hands = append(hands, r)
	}

	if err != io.EOF {
		fmt.Println(err)
		return
	}

	calcWinnings(hands, false)
	calcWinnings(hands, true)
}

func calcWinnings(hands []round, partTwo bool) {
	sort.Slice(hands, func(i, j int) bool {
		return !comapre(hands[i], hands[j], partTwo)
	})

	var winnings int

	for i, r := range hands {
		winnings += (i + 1) * r.Bet
	}

	fmt.Println(winnings)
}

// return true if a is better than b
func comapre(a round, b round, partTwo bool) bool {
	if partTwo && a.partTwo != b.partTwo {
		return a.partTwo > b.partTwo
	}
	if !partTwo && a.Strength != b.Strength {
		return a.Strength > b.Strength
	}

	for i := 0; i < 5; i++ {
		aValue := getCardValue(rune(a.Hand[i]), partTwo)
		bValue := getCardValue(rune(b.Hand[i]), partTwo)
		if aValue != bValue {
			return aValue > bValue
		}
	}

	return false
}

// calculate hand strength
func calcStrength(hand string, partTwo bool) strength {
	chars := make(map[rune]int, 5)
	for _, card := range hand {
		chars[card] += 1
	}

	countA := 0
	countB := 0

	for k, v := range chars {
		if partTwo && k == 'J' {
			continue
		}
		if v >= countA {
			countB = countA
			countA = v
		} else if v > countB {
			countB = v
		}
	}

	if partTwo {
		countA += chars['J']
	}

	// 5 same
	if countA == 5 {
		return 7
	}
	// poker
	if countA == 4 {
		return 6
	}
	// full house
	if countA == 3 && countB == 2 {
		return 5
	}
	// drill
	if countA == 3 {
		return 4
	}
	// two pairs
	if countA == 2 && countB == 2 {
		return 3
	}
	// pair
	if countA == 2 {
		return 2
	}
	return 1
}

func getCardValue(card rune, partTwo bool) byte {
	switch card {
	case 'J':
		if partTwo {
			return 1
		}
		return 11
	case 'T':
		return 10
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return byte(card) - '0'
	}
}
