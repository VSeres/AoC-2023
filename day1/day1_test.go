package day1

import (
	"bufio"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	file, err := os.Open("test_input_1.txt")
	if err != nil {
		t.Fatalf("Faild to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += partOne(line)
	}

	const result = 142
	if sum != result {
		t.Errorf("Invalid sum %d should be %d", sum, result)
	}
}

var lineValues = []int{
	29,
	83,
	13,
	24,
	42,
	14,
	76,
	38,
	11,
	11,
}

func TestPartTwo(t *testing.T) {
	file, err := os.Open("test_input_2.txt")
	if err != nil {
		t.Fatalf("Faild to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		num := partTwo(line)
		if lineValues[i] != num {
			t.Errorf("Invalid sum %d should be %d", num, lineValues[i])
		}
		i++
		sum += num
	}

	const result = 281 + 38 + 11 + 11
	if sum != result {
		t.Errorf("Invalid sum %d should be %d", sum, result)
	}
}
