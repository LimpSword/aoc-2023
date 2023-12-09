package main

import (
	"regexp"
	"strings"
)

type Next struct {
	Left  string
	Right string
}

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	moveOrder := make([]string, 0)

	goTo := make(map[string]Next)

	for k, line := range lines {
		if line == "" {
			continue
		}
		if k == 0 {
			for _, c := range line {
				moveOrder = append(moveOrder, string(c))
			}
		} else {
			regex := regexp.MustCompile("([A-Z]{3}) = \\(([A-Z]{3}), ([A-Z]{3})\\)")
			matches := regex.FindStringSubmatch(line)
			moveFrom := matches[1]
			moveToLeft := matches[2]
			moveToRight := matches[3]

			goTo[moveFrom] = Next{moveToLeft, moveToRight}
		}
	}

	current := "AAA"
	k := 0
	for current != "ZZZ" {
		left := moveOrder[k%len(moveOrder)] == "L"

		next := goTo[current]
		if left {
			current = next.Left
		} else {
			current = next.Right
		}
		k++
	}

	return k
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	moveOrder := make([]string, 0)

	goTo := make(map[string]Next)

	for k, line := range lines {
		if line == "" {
			continue
		}
		if k == 0 {
			for _, c := range line {
				moveOrder = append(moveOrder, string(c))
			}
		} else {
			regex := regexp.MustCompile("([A-Z1-9]{3}) = \\(([A-Z1-9]{3}), ([A-Z1-9]{3})\\)")
			matches := regex.FindStringSubmatch(line)
			moveFrom := matches[1]
			moveToLeft := matches[2]
			moveToRight := matches[3]

			goTo[moveFrom] = Next{moveToLeft, moveToRight}
		}
	}

	starting := make([]string, 0)
	for k, _ := range goTo {
		if strings.HasSuffix(k, "A") {
			starting = append(starting, k)
		}
	}

	var current []string
	kMap := make(map[string]int)
	for _, s := range starting {
		current = append(current, s)
		kMap[s] = 0
	}
	for c := range current {
		for !strings.HasSuffix(current[c], "Z") {
			left := moveOrder[kMap[starting[c]]%len(moveOrder)] == "L"

			next := goTo[current[c]]
			if left {
				current[c] = next.Left
			} else {
				current[c] = next.Right
			}
			kMap[starting[c]] += 1
		}
	}

	periods := make([]int, 0)
	for _, k := range kMap {
		periods = append(periods, k)
	}

	return lcm(periods[0], periods[1], periods[2:]...)
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}
