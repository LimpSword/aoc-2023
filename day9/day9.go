package main

import (
	"slices"
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
		for !slices.Equal(slices.Compact(slices.Clone(v[len(v)-1])), []int{0}) {
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
		v[len(v)-1] = append(v[len(v)-1], 0)
		for i := len(v) - 2; i >= 0; i-- {
			v[i] = append(v[i], v[i+1][len(v[i+1])-1]+v[i][len(v[i])-1])
		}
		sum += v[0][len(v[0])-1]
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
		for !slices.Equal(slices.Compact(slices.Clone(v[len(v)-1])), []int{0}) {
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
		for i := len(v) - 2; i >= 0; i-- {
			v[i] = append([]int{-v[i+1][0] + v[i][0]}, v[i]...)
		}
		sum += v[0][0]
	}

	return sum
}
