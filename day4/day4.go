package main

import (
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	sum := 0
	for _, line := range lines {
		line := strings.Split(line, ": ")[1]
		splt := strings.Split(line, " | ")

		winning := splt[0]
		myNumbers := splt[1]

		winningSplt := strings.Split(winning, " ")
		myNumbersSplt := strings.Split(myNumbers, " ")

		s := 0

		for _, num := range myNumbersSplt {
			for _, win := range winningSplt {
				if num == win && num != "" && win != "" {
					if s == 0 {
						s = 1
					} else {
						s *= 2
					}
				}
			}
		}
		sum += s
	}

	return sum
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	added := make(map[int]int)
	addedCount := len(lines)

	for k, line := range lines {
		line := strings.Split(line, ": ")[1]
		splt := strings.Split(line, " | ")

		winning := splt[0]
		myNumbers := splt[1]

		winningSplt := strings.Split(winning, " ")
		myNumbersSplt := strings.Split(myNumbers, " ")

		s := 0

		for _, num := range myNumbersSplt {
			for _, win := range winningSplt {
				if num == win && num != "" && win != "" {
					s++
				}
			}
		}
		for i := k + 1; i < k+s+1; i++ {
			if added[k+1] == 0 {
				added[i+1] = 1
				addedCount++
			} else {
				added[i+1] += 1 + added[k+1]
				addedCount += 1 + added[k+1]
			}
		}
	}

	return addedCount
}
