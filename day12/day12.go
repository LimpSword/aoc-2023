package main

import (
	"strconv"
	"strings"
)

var cache = make(map[string]int)

func evaluate(line string, values []int) int {
	key := line
	for _, v := range values {
		key += strconv.Itoa(v)
	}
	if cached, ok := cache[key]; ok {
		return cached
	}

	if line == "" {
		if len(values) == 0 {
			return 1
		}
		return 0
	}
	if len(values) == 0 {
		if strings.ContainsRune(line, '#') {
			return 0
		}
		// everything can be a .
		return 1
	}

	count := 0

	if line[0] == '.' || line[0] == '?' {
		// continue with the same values
		count += evaluate(line[1:], values)
	}

	if line[0] == '#' || line[0] == '?' {
		if values[0] <= len(line) {
			// is line long enough to place an # ?
			if !strings.ContainsRune(line[:values[0]], '.') {
				// are the next values all # or ?
				if values[0] == len(line) || line[values[0]] != '#' {
					// place an # if it is the end of the line or the next char is not an #
					if values[0] == len(line) {
						count += evaluate(line[values[0]:], values[1:])
					} else {
						count += evaluate(line[values[0]+1:], values[1:])
					}
				}
			}
		}
	}
	cache[key] = count

	return count
}

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	p1 := 0
	for _, line := range lines {
		splt := strings.Split(line, " ")
		l := splt[0]

		values := make([]int, 0)
		for _, v := range strings.Split(splt[1], ",") {
			j, _ := strconv.Atoi(v)
			values = append(values, j)
		}

		p1 += evaluate(l, values)
	}

	return p1
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	lines := strings.Split(text, "\n")

	p2 := 0
	for _, line := range lines {
		splt := strings.Split(line, " ")
		l := splt[0]
		lFinal := ""

		for i := 0; i < 5; i++ {
			lFinal += l
			if i < 4 {
				lFinal += "?"
			}
		}

		values := make([]int, 0)
		for _, v := range strings.Split(splt[1], ",") {
			j, _ := strconv.Atoi(v)
			values = append(values, j)
		}
		valuesFinal := make([]int, 0)
		for i := 0; i < 5; i++ {
			valuesFinal = append(valuesFinal, values...)
		}

		p2 += evaluate(lFinal, valuesFinal)
	}
	return p2
}
