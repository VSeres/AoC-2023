package day12

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func Solve() {
	file, err := os.Open("day12/input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		spring     []byte
		groupsData []byte
	)

	sum := 0

	for {
		_, err = fmt.Fscanln(file, &spring, &groupsData)
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		groups := make([]int, 0, 32)
		num := 0
		total := 0
		for _, v := range groupsData {
			if v != ',' {
				num = num*10 + int(v-'0')
			} else {
				groups = append(groups, num)
				total += num
				num = 0
			}
		}
		groups = append(groups, num)
		total += num

		fmt.Println("---")
		c := find(spring, groups, 0, total, "")
		sum += c
		fmt.Println(string(spring), string(groupsData), "  ", c)
	}
	fmt.Println(sum)
}

func countSprings(list []byte, offset int) int {
	count := 0
	if offset >= len(list) {
		return 0
	}
	for _, v := range list[offset:] {
		if v == '#' {
			count++
		}
	}

	return count
}

func find(spring []byte, group []int, i int, total int, option string) int {
	if countSprings(spring, i) > 0 && len(group) == 0 { // no more groups, but there are still springs
		return 0
	}

	if len(group) > 0 && len(spring) <= i { // no more springs, but there are still groups
		return 0
	}

	if len(group) == 0 {

		if strings.Count(option, "#") > total {
			return 0
		}

		fmt.Println(option)

		return 1
	}

	skipped := ""
	// next valid position
	for spring[i] != '?' && spring[i] != '#' {
		i += 1
		skipped += "."
		if i >= len(spring) {
			return 0
		}
	}

	g := group[0]
	// inster group
	if len(spring) < i+g { // cannot fit group
		return 0
	}

	fits := true
	for j := i; j < i+g; j++ {
		if spring[j] == '.' {
			fits = false
			break
		}
	}

	lookAhead := len(spring) <= i+g || spring[i+g] != '#'
	lookBehind := i-1 < 0 || spring[i-1] != '#' // alway true if in the previous run a group was used
	res := 0
	if fits && lookAhead && lookBehind {
		a := ""
		if len(spring) > i+g {
			a = string(spring[i+g])
		}
		res = find(spring, group[1:], i+g+1, total, option+skipped+strings.Repeat("#", g)+a)
	}

	res += find(spring, group, i+1, total, option+skipped+string(spring[i]))

	return res
}
