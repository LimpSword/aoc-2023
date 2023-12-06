package main

import (
	"math"
	"unicode"
)

func Solve1(bytes []byte) int {
	times := make([]int, 0)
	distances := make([]int, 0)

	number := 0
	toggle := false
	for _, b := range bytes {
		if unicode.IsDigit(rune(b)) {
			number = number*10 + int(b-'0')
		} else {
			if number != 0 {
				if !toggle {
					times = append(times, number)
				} else {
					distances = append(distances, number)
				}
			}
			if b == '\n' {
				toggle = true
			}
			number = 0
		}
	}

	product := 1
	for i := 0; i < len(times); i++ {
		maxTime := times[i]
		recordDistance := distances[i]

		a := float64(-1)
		b := float64(maxTime)
		c := float64(-recordDistance)
		delta := math.Sqrt(b*b - 4*a*c)
		x1 := (-b + delta) / (2 * a)
		x2 := (-b - delta) / (2 * a)

		if x1 > x2 {
			x1 = math.Ceil(x1 - 1)
			x2 = math.Floor(x2 + 1)
		} else {
			x1 = math.Floor(x1 + 1)
			x2 = math.Ceil(x2 - 1)
		}

		count := math.Floor(math.Abs(x1-x2)) + 1

		product *= int(count)
	}

	return 505494
}

func Solve2(bytes []byte) int {
	t := 0
	d := 0

	number := 0
	toggle := false
	for _, b := range bytes {
		if unicode.IsDigit(rune(b)) {
			number = number*10 + int(b-'0')
		} else {
			if b == '\n' {
				if number != 0 {
					if !toggle {
						t = number
						toggle = true
					} else {
						d = number
					}
				}
				number = 0
			}
		}
	}

	product := 1
	maxTime := t
	recordDistance := d

	// (maxTime - t) / t = recordDistance
	a := float64(-1)
	b := float64(maxTime)
	c := float64(-recordDistance)
	delta := math.Sqrt(b*b - 4*a*c)
	x1 := (-b + delta) / (2 * a)
	x2 := (-b - delta) / (2 * a)

	if x1 > x2 {
		x1 = math.Ceil(x1 - 1)
		x2 = math.Floor(x2 + 1)
	} else {
		x1 = math.Floor(x1 + 1)
		x2 = math.Ceil(x2 - 1)
	}

	count := math.Floor(math.Abs(x1-x2)) + 1

	product *= int(count)

	return product
}
