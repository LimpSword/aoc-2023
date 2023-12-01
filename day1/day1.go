package main

import (
	"os"
	"strings"
	"unicode"
)

func Solve1() int {
	bytes, _ := os.ReadFile("input.txt")
	input := string(bytes)

	// Part 1
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		first, last := -1, -1
		for k := range line {
			if unicode.IsDigit(rune(line[k])) {
				first = int(line[k] - '0')
				break
			}
		}
		for k := len(line) - 1; k >= 0; k-- {
			if unicode.IsDigit(rune(line[k])) {
				last = int(line[k] - '0')
				break
			}
		}
		sum += 10*first + last
	}
	return sum
}

func Solve2() int {
	bytes, _ := os.ReadFile("input.txt")
	input := string(bytes)

	// Part 1
	lines := strings.Split(input, "\n")
	sum := 0

	numbers := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	for _, line := range lines {
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if first == -1 {
					first = int(line[i] - '0')
				} else {
					last = int(line[i] - '0')
				}
			} else {
				for j := i; j < min(len(line), i+6); j++ {
					if n, ok := numbers[line[i:j+1]]; ok {
						if first == -1 {
							first = n
							i = j
						} else {
							last = n
							i = j
						}
					}
				}
			}
		}
		if last == -1 {
			sum += 10*first + first
		} else {
			sum += 10*first + last
		}
	}
	return sum
}
