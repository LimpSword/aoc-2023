package main

import (
	"strings"
	"unicode"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	sum := 0
	for i, line := range lines {
		number := 0
		startCol := -1
		for j, char := range line {
			if unicode.IsDigit(char) {
				if startCol == -1 {
					startCol = j
				}
				number = number*10 + int(char-'0')
			} else {
				// check if the number is a part number
				if number != 0 {
					if isPartNumber(lines, i, startCol, j-1) {
						sum += number
					}
					number = 0
					startCol = -1
				}
			}
		}
		if number != 0 {
			if isPartNumber(lines, i, startCol, len(line)-1) {
				sum += number
			}
			number = 0
			startCol = -1
		}
	}
	return sum
}

func isPartNumber(lines []string, line int, startCol int, column int) bool {
	for i := max(line-1, 0); i < min(line+2, len(lines)); i++ {
		for j := max(startCol-1, 0); j < min(column+2, len(lines[0])); j++ {
			if lines[i][j] != '.' && !unicode.IsDigit(rune(lines[i][j])) {
				return true
			}
		}
	}
	return false
}

func isPartNumber2(gears map[int][]int, number int, lines []string, line int, startCol int, column int) bool {
	found := false
	for i := max(line-1, 0); i < min(line+2, len(lines)); i++ {
		for j := max(startCol-1, 0); j < min(column+2, len(lines[0])); j++ {
			if lines[i][j] == '*' {
				found = true
				c := cantor(i, j)
				if _, ok := gears[c]; !ok {
					gears[cantor(i, j)] = make([]int, 0)
				}
				gears[c] = append(gears[c], number)
			}
		}
	}
	return found
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	gears := make(map[int][]int)

	sum := 0
	for i, line := range lines {
		number := 0
		startCol := -1
		for j, char := range line {
			if unicode.IsDigit(char) {
				if startCol == -1 {
					startCol = j
				}
				number = number*10 + int(char-'0')
			} else {
				// check if the number is a part number
				if number != 0 {
					if isPartNumber2(gears, number, lines, i, startCol, j-1) {
						sum += number
					}
					number = 0
					startCol = -1
				}
			}
		}
		if number != 0 {
			if isPartNumber(lines, i, startCol, len(line)-1) {
				sum += number
			}
			number = 0
			startCol = -1
		}
	}

	gearsSum := 0
	for _, v := range gears {
		if len(v) == 2 {
			gearsSum += v[0] * v[1]
		}
	}

	return gearsSum
}

func cantor(a int, b int) int {
	return (a+b)*(a+b+1)/2 + b
}
