package main

import (
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Solve1() int {
	bytes, _ := os.ReadFile("day1/input.txt")
	input := string(bytes)

	// Part 1
	lines := strings.Split(input, "\n")
	sum := 0
	for _, line := range lines {
		first := -1
		last := -1
		for _, c := range line {
			if first == -1 && unicode.IsDigit(c) {
				first = int(c - '0')
			} else {
				if unicode.IsDigit(c) {
					last = int(c - '0')
				}
			}
		}
		if last == -1 {
			a, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(first))
			sum += a
		} else {
			a, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
			sum += a
		}
	}
	return sum
}

func Solve2() int {
	bytes, _ := os.ReadFile("day1/input.txt")
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
		first := -1
		last := -1
		for i := 0; i < len(line); i++ {
			for j := i; j < len(line); j++ {
				if n, ok := numbers[line[i:j]]; ok {
					if first == -1 {
						first = n
						i = j - 1
					} else {
						last = n
						i = j - 1
					}
				} else if math.Abs(float64(i-j)) == 0 {
					if unicode.IsDigit(rune(line[i])) {
						if first == -1 {
							first = int(line[i] - '0')
						} else {
							last = int(line[i] - '0')
						}
					}
				}
			}
		}
		if last == -1 {
			a, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(first))
			sum += a
		} else {
			a, _ := strconv.Atoi(strconv.Itoa(first) + strconv.Itoa(last))
			sum += a
		}
	}
	return sum
}
