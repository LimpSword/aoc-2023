package main

import (
	"slices"
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	emptyLines := make([]int, 0)
	emptyColumns := make([]int, 0)

	// determine empty lines indexes
	for i := range lines {
		empty := true
		for j := range lines[i] {
			if lines[i][j] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyLines = append(emptyLines, i)
		}
	}
	// determine empty columns indexes
	for i := range lines[0] {
		empty := true
		for j := range lines {
			if lines[j][i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	// determine each pair of points
	points := make([][]int, 0)
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == '#' {
				points = append(points, []int{i, j})
			}
		}
	}

	// distance between each pair of points
	total := 0
	for p1 := 0; p1 < len(points); p1++ {
		for p2 := p1 + 1; p2 < len(points); p2++ {
			dist := 0 // this is the Manhattan distance (count of lines and columns)
			for i := min(points[p1][0], points[p2][0]); i < max(points[p1][0], points[p2][0]); i++ {
				if slices.Contains(emptyLines, i) {
					dist++
				}
				dist++
			}
			for i := min(points[p1][1], points[p2][1]); i < max(points[p1][1], points[p2][1]); i++ {
				if slices.Contains(emptyColumns, i) {
					dist++
				}
				dist++
			}
			total += dist
		}
	}

	return total
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	emptyLines := make([]int, 0)
	emptyColumns := make([]int, 0)

	for i := range lines {
		empty := true
		for j := range lines[i] {
			if lines[i][j] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyLines = append(emptyLines, i)
		}
	}
	for i := range lines[0] {
		empty := true
		for j := range lines {
			if lines[j][i] == '#' {
				empty = false
				break
			}
		}
		if empty {
			emptyColumns = append(emptyColumns, i)
		}
	}

	points := make([][]int, 0)
	for i, _ := range lines {
		for j, _ := range lines[i] {
			if lines[i][j] == '#' {
				points = append(points, []int{i, j})
			}
		}
	}

	total := 0
	for p1 := 0; p1 < len(points); p1++ {
		for p2 := p1 + 1; p2 < len(points); p2++ {
			dist := 0
			for i := min(points[p1][0], points[p2][0]); i < max(points[p1][0], points[p2][0]); i++ {
				if slices.Contains(emptyLines, i) {
					// replace 1 by 1000000
					dist += 999999
				}
				dist++
			}
			for i := min(points[p1][1], points[p2][1]); i < max(points[p1][1], points[p2][1]); i++ {
				if slices.Contains(emptyColumns, i) {
					// replace 1 by 1000000
					dist += 999999
				}
				dist++
			}
			total += dist
		}
	}

	return total
}
