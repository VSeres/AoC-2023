package day12

import (
	"fmt"
	"io"
	"os"
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
		for _, v := range groupsData {
			if v != ',' {
				groups = append(groups, int(v-'0'))
			}

		}
		fmt.Println("---")
		c := find(spring, groups, 0)
		sum += c
		fmt.Println("hit:", hit, "lookup:", lookup)
		hit = 0
		lookup = 0
		fmt.Println(string(spring), string(groupsData), "  ", c)
		for k := range cache {
			delete(cache, k)
		}
	}
	fmt.Println(sum)
}

type key struct {
	g int
	i int
}

var cache = make(map[key]int, 0)
var (
	lookup = 0
	hit    = 0
)

func find(springs []byte, groups []int, i int) int {
	key := key{g: len(groups), i: i}
	lookup++
	if v, ok := cache[key]; ok {
		hit++
		return v
	}

	if len(groups) == 0 && len(springs) >= i && countSprings(springs[i:]) > 0 { // no more groups, but there are still springs
		return 0
	} else if len(groups) == 0 { // no more groups and springs
		return 1
	} else if groups[0] >= len(springs) { // next group is bigger than the length
		return 0
	}

	if i >= len(springs) && len(groups) > 0 { // no more groups, but there are still springs ??????
		return 0
	}

	for springs[i] != '#' && springs[i] != '?' {
		i++
		if len(springs) <= i {
			return 0
		}
	}

	g := groups[0]
	if g+i > len(springs) { // not enough space for group
		return 0
	}
	fits := true
	for j := i; j < i+g; j++ {
		if springs[j] == '.' {
			// cannot fit
			fits = false
			break
		}
	}
	lookahead := i+g >= len(springs) || springs[i+g] == '?' || springs[i+g] == '.'
	var res int
	if fits && lookahead {
		res = find(springs, groups[1:], i+g+1)
	}

	if springs[i] == '?' {
		res += find(springs, groups, i+1)
	}
	cache[key] = res
	return res
}

func countSprings(list []byte) int {
	count := 0
	for _, v := range list {
		if v == '#' {
			count++
		}
	}
	return count

}
