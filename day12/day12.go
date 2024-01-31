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

	sumPartOne := 0
	sumPartTwo := 0

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

		//part one
		// sumPartOne += find(spring, groups, 0, total, "")
		evictCache()
		// part two
		part := partTwo(spring, groups, total)
		sumPartTwo += part
		evictCache()
	}
	fmt.Println(sumPartOne, sumPartTwo)
}

func evictCache() {
	for ck := range cache {
		delete(cache, ck)
	}
}

func partTwo(spring []byte, groups []int, total int) int {
	unfolded := make([]byte, 0, len(spring)*5+4)
	unfolded = append(unfolded, spring...)
	unfoldedGroup := make([]int, 0, len(groups)*5)
	for i := 0; i < 4; i++ {
		unfolded = append(unfolded, '?')
		unfolded = append(unfolded, spring...)
		unfoldedGroup = append(unfoldedGroup, groups...)

	}
	unfoldedGroup = append(unfoldedGroup, groups...)
	fmt.Printf("%s %v\n", string(unfolded), unfoldedGroup)
	return find(unfolded, unfoldedGroup, 0, total*5, "")
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

type cKey struct {
	a int
	b int
}

var cache = make(map[string]int, 0)

func find(spring []byte, group []int, i int, total int, option string) int {
	// key := cKey{
	// 	a: len(group),
	// 	b: i,
	// }
	if v, ok := cache[option]; ok {
		return v
	}
	if countSprings(spring, i) > 0 && len(group) == 0 { // no more groups, but there are still springs
		return 0
	}

	if len(group) > 0 && len(spring) <= i { // no more springs, but there are still groups
		return 0
	}

	if len(group) == 0 {
		c := strings.Count(option, "#")
		if c > total {
			return 0
		}
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
	cache[option] = res
	return res
}
