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

func Solve2Mad() int {
	bytes, _ := os.ReadFile("input.txt")
	input := string(bytes)

	// Part 2
	lines := strings.Split(input, "\n")
	sum := 0

	for _, line := range lines {
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = int(line[i] - '0')
				break
			} else {
				if i+2 < len(line) && line[i:i+3] == "one" {
					first = 1
					i++
					break
				} else if i+2 < len(line) && line[i:i+3] == "two" {
					first = 2
					i += 2
					break
				} else if i+4 < len(line) && line[i:i+5] == "three" {
					first = 3
					i += 4
					break
				} else if i+3 < len(line) && line[i:i+4] == "four" {
					first = 4
					i += 3
					break
				} else if i+3 < len(line) && line[i:i+4] == "five" {
					first = 5
					i += 3
					break
				} else if i+2 < len(line) && line[i:i+3] == "six" {
					first = 6
					i += 2
					break
				} else if i+4 < len(line) && line[i:i+5] == "seven" {
					first = 7
					i += 4
					break
				} else if i+4 < len(line) && line[i:i+5] == "eight" {
					first = 8
					i += 4
					break
				} else if i+3 < len(line) && line[i:i+4] == "nine" {
					first = 9
					i += 3
					break
				}
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = int(line[i] - '0')
				break
			} else {
				if i >= 2 && line[i-2:i+1] == "one" {
					last = 1
					break
				} else if i >= 2 && line[i-2:i+1] == "two" {
					last = 2
					break
				} else if i >= 4 && line[i-4:i+1] == "three" {
					last = 3
					break
				} else if i >= 3 && line[i-3:i+1] == "four" {
					last = 4
					break
				} else if i >= 3 && line[i-3:i+1] == "five" {
					last = 5
					break
				} else if i >= 2 && line[i-2:i+1] == "six" {
					last = 6
					break
				} else if i >= 4 && line[i-4:i+1] == "seven" {
					last = 7
					break
				} else if i >= 4 && line[i-4:i+1] == "eight" {
					last = 8
					break
				} else if i >= 3 && line[i-3:i+1] == "nine" {
					last = 9
					break
				}
			}
		}
		sum += 10*first + last
	}
	return sum
}
