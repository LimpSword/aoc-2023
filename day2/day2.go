package main

import (
	"bufio"
	bytes2 "bytes"
	"regexp"
	"strconv"
	"strings"
)

func Solve1(bytes []byte) int {
	buffer := bufio.NewReader(bytes2.NewBuffer(bytes))

	// Part 1
	valid_games := 0
	k := 1
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.Split(line, ":")[1]

		blueRegex := regexp.MustCompile(`(\d+) blue`)
		redRegex := regexp.MustCompile(`(\d+) red`)
		greenRegex := regexp.MustCompile(`(\d+) green`)

		// get all integers from the groups of blue regex
		blueMatches := blueRegex.FindAllStringSubmatch(line, -1)
		blue := 0
		for _, match := range blueMatches {
			b, _ := strconv.Atoi(match[1])
			blue = max(blue, b)
		}

		// get all integers from the groups of red regex
		redMatches := redRegex.FindAllStringSubmatch(line, -1)
		red := 0
		for _, match := range redMatches {
			r, _ := strconv.Atoi(match[1])
			red = max(red, r)
		}

		// get all integers from the groups of green regex
		greenMatches := greenRegex.FindAllStringSubmatch(line, -1)
		green := 0
		for _, match := range greenMatches {
			g, _ := strconv.Atoi(match[1])
			green = max(green, g)
		}

		if red <= 12 && green <= 13 && blue <= 14 {
			valid_games += k
		}
		k++
	}
	return valid_games
}

func Solve2(bytes []byte) int {
	buffer := bufio.NewReader(bytes2.NewBuffer(bytes))

	// Part 1
	minimum_cubes := 0
	k := 1
	for {
		line, err := buffer.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.Split(line, ":")[1]

		blueRegex := regexp.MustCompile(`(\d+) blue`)
		redRegex := regexp.MustCompile(`(\d+) red`)
		greenRegex := regexp.MustCompile(`(\d+) green`)

		// get all integers from the groups of blue regex
		blueMatches := blueRegex.FindAllStringSubmatch(line, -1)
		blue := 0
		for _, match := range blueMatches {
			b, _ := strconv.Atoi(match[1])
			blue = max(blue, b)
		}

		// get all integers from the groups of red regex
		redMatches := redRegex.FindAllStringSubmatch(line, -1)
		red := 0
		for _, match := range redMatches {
			r, _ := strconv.Atoi(match[1])
			red = max(red, r)
		}

		// get all integers from the groups of green regex
		greenMatches := greenRegex.FindAllStringSubmatch(line, -1)
		green := 0
		for _, match := range greenMatches {
			g, _ := strconv.Atoi(match[1])
			green = max(green, g)
		}

		minimum_cubes += red * green * blue
		k++
	}
	return minimum_cubes
}
