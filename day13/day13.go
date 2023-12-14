package main

import (
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	patterns := strings.Split(text, "\n\n")

	p1 := 0
	for _, pattern := range patterns {
		grid := strings.Split(pattern, "\n")

		found := false

		// replace 1 letter of
		for i := 1; i < len(grid); i++ {
			// are the two sides equal?
			eq := true
			minLines := min(len(grid)-i, i)
			for k := 0; k < minLines; k++ {
				for j := 0; j < len(grid[i]); j++ {
					if grid[i+k][j] != grid[i-k-1][j] {
						eq = false
						break
					}
				}
			}
			if eq {
				found = true
				p1 += 100 * i
				break
			}
		}

		if !found {
			// check columns
			for i := 1; i < len(grid[0]); i++ {
				// are the two sides equal?
				eq := true
				minLines := min(len(grid[0])-i, i)
				for k := 0; k < minLines; k++ {
					for j := 0; j < len(grid); j++ {
						if grid[j][i+k] != grid[j][i-k-1] {
							eq = false
							break
						}
					}
				}
				if eq {
					found = true
					p1 += i
					break
				}
			}
		}
	}
	return p1
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	patterns := strings.Split(text, "\n\n")

	p1 := 0
	for _, pattern := range patterns {
		grid := strings.Split(pattern, "\n")

		found := false

		// replace 1 letter of
		for i := 1; i < len(grid); i++ {
			// are the two sides equal?
			eq := true
			minLines := min(len(grid)-i, i)
			fixedSmudge := false
			for k := 0; k < minLines; k++ {
				for j := 0; j < len(grid[i]); j++ {
					if grid[i+k][j] != grid[i-k-1][j] {
						if fixedSmudge {
							eq = false
							break
						}
						fixedSmudge = true
					}
				}
			}
			if eq && fixedSmudge {
				found = true
				p1 += 100 * i
				break
			}
		}

		if !found {
			// check columns
			for i := 1; i < len(grid[0]); i++ {
				// are the two sides equal?
				eq := true
				minLines := min(len(grid[0])-i, i)
				fixedSmudge := false
				for k := 0; k < minLines; k++ {
					for j := 0; j < len(grid); j++ {
						if grid[j][i+k] != grid[j][i-k-1] {
							if fixedSmudge {
								eq = false
								break
							}
							fixedSmudge = true
						}
					}
				}
				if eq && fixedSmudge {
					found = true
					p1 += i
					break
				}
			}
		}
	}
	return p1
}
