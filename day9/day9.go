package main

import (
	"strconv"
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")

	sum := 0

	successive := make(map[int][][]int)
	for k, line := range lines {
		first := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(n)
			first = append(first, number)
		}
		successive[k] = [][]int{first}
	}

	for k, v := range successive {
		for !isFullZero(v[len(v)-1]) {
			last := v[len(v)-1]
			next := make([]int, len(last)-1)
			for i := 0; i < len(last)-1; i++ {
				next[i] = last[i+1] - last[i]
			}
			v = append(v, next)
		}
		successive[k] = v
	}

	for _, v := range successive {
		r := 0
		for i := len(v) - 2; i >= 0; i-- {
			r = v[i][len(v[i])-1] + r
		}
		sum += r
	}

	return sum
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")

	sum := 0

	successive := make(map[int][][]int)
	for k, line := range lines {
		first := make([]int, 0)
		for _, n := range strings.Split(line, " ") {
			number, _ := strconv.Atoi(n)
			first = append(first, number)
		}
		successive[k] = [][]int{first}
	}

	for k, v := range successive {
		for !isFullZero(v[len(v)-1]) {
			last := v[len(v)-1]
			next := make([]int, len(last)-1)
			for i := 0; i < len(last)-1; i++ {
				next[i] = last[i+1] - last[i]
			}
			v = append(v, next)
		}
		successive[k] = v
	}

	for _, v := range successive {
		r := 0
		for i := len(v) - 2; i >= 0; i-- {
			r = v[i][0] - r
		}
		sum += r
	}

	return sum
}

func isFullZero(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}
