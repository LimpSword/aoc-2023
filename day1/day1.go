package main

import (
	"bufio"
	bytes2 "bytes"
	"os"
	"strings"
	"unicode"
)

func Solve1() int {
	bytes, _ := os.ReadFile("input.txt")

	buffer := bufio.NewReader(bytes2.NewBuffer(bytes))

	// Part 1

	sum := 0
	for {
		line, _, err := buffer.ReadLine()
		if err != nil {
			break
		}
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

	// Part 2
	buffer := bufio.NewReader(bytes2.NewBuffer(bytes))
	sum := 0

	for {
		line, _, err := buffer.ReadLine()
		if err != nil {
			break
		}
		first, last := -1, -1
		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				first = int(line[i] - '0')
				break
			} else if i+2 < len(line) {
				if rune(line[i]) == 'o' && rune(line[i+1]) == 'n' && rune(line[i+2]) == 'e' {
					first = 1
					break
				} else if rune(line[i]) == 't' && rune(line[i+1]) == 'w' && rune(line[i+2]) == 'o' {
					first = 2
					break
				} else if rune(line[i]) == 's' && rune(line[i+1]) == 'i' && rune(line[i+2]) == 'x' {
					first = 6
					break
				} else if i+3 < len(line) {
					if rune(line[i]) == 'f' {
						if rune(line[i+1]) == 'o' && rune(line[i+2]) == 'u' && rune(line[i+3]) == 'r' {
							first = 4
							break
						} else if rune(line[i+1]) == 'i' && rune(line[i+2]) == 'v' && rune(line[i+3]) == 'e' {
							first = 5
							break
						}
					} else if rune(line[i]) == 'n' && rune(line[i+1]) == 'i' && rune(line[i+2]) == 'n' && rune(line[i+3]) == 'e' {
						first = 9
						break
					} else if i+4 < len(line) {
						if rune(line[i]) == 't' && rune(line[i+1]) == 'h' && rune(line[i+2]) == 'r' && rune(line[i+3]) == 'e' && rune(line[i+4]) == 'e' {
							first = 3
							break
						} else if rune(line[i]) == 's' && rune(line[i+1]) == 'e' && rune(line[i+2]) == 'v' && rune(line[i+3]) == 'e' && rune(line[i+4]) == 'n' {
							first = 7
							break
						} else if rune(line[i]) == 'e' && rune(line[i+1]) == 'i' && rune(line[i+2]) == 'g' && rune(line[i+3]) == 'h' && rune(line[i+4]) == 't' {
							first = 8
							break
						}
					}
				}
			}
		}
		for i := len(line) - 1; i >= 0; i-- {
			if unicode.IsDigit(rune(line[i])) {
				last = int(line[i] - '0')
				break
			} else if i >= 2 {
				if rune(line[i-2]) == 'o' && rune(line[i-1]) == 'n' && rune(line[i]) == 'e' {
					last = 1
					break
				} else if rune(line[i-2]) == 't' && rune(line[i-1]) == 'w' && rune(line[i]) == 'o' {
					last = 2
					break
				} else if rune(line[i-2]) == 's' && rune(line[i-1]) == 'i' && rune(line[i]) == 'x' {
					last = 6
					break
				} else if i >= 3 {
					if rune(line[i-3]) == 'f' {
						if rune(line[i-2]) == 'o' && rune(line[i-1]) == 'u' && rune(line[i]) == 'r' {
							last = 4
							break
						} else if rune(line[i-2]) == 'i' && rune(line[i-1]) == 'v' && rune(line[i]) == 'e' {
							last = 5
							break
						}
					} else if rune(line[i-3]) == 'n' && rune(line[i-2]) == 'i' && rune(line[i-1]) == 'n' && rune(line[i]) == 'e' {
						last = 9
						break
					} else if i >= 4 {
						if rune(line[i-4]) == 't' && rune(line[i-3]) == 'h' && rune(line[i-2]) == 'r' && rune(line[i-1]) == 'e' && rune(line[i]) == 'e' {
							last = 3
							break
						} else if rune(line[i-4]) == 's' && rune(line[i-3]) == 'e' && rune(line[i-2]) == 'v' && rune(line[i-1]) == 'e' && rune(line[i]) == 'n' {
							last = 7
							break
						} else if rune(line[i-4]) == 'e' && rune(line[i-3]) == 'i' && rune(line[i-2]) == 'g' && rune(line[i-1]) == 'h' && rune(line[i]) == 't' {
							last = 8
							break
						}
					}
				}
			}
		}
		sum += 10*first + last
	}
	return sum
}
