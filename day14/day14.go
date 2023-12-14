package main

import (
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	array := strings.Split(text, "\n")

	result := 0
	columnLength := len(array)

	bestPerColumn := make([]int, len(array[0]))
	// for each column
	for i := 0; i < len(array[0]); i++ {
		bestPerColumn[i] = columnLength

		for j := 0; j < columnLength; j++ {
			if array[j][i] == '#' {
				bestPerColumn[i] = columnLength - j - 1
			} else if array[j][i] == 'O' {
				result += bestPerColumn[i]
				bestPerColumn[i]--
			}
		}
	}

	return result
}

func Solve2(bytes []byte) int {
	return 0
}
