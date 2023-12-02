package main

import (
	"bufio"
	bytes2 "bytes"
	"unicode"
)

func Solve2Mad(bytes []byte) int {
	reader := bufio.NewReader(bytes2.NewBuffer(bytes))

	red := []byte("red")
	green := []byte("green")
	blue := []byte("blue")

	valid_games := 0
	k := 1
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		line = bytes2.Split(line, []byte(":"))[1]
		sets := bytes2.Split(line, []byte(";"))

		r, g, b := 0, 0, 0
		for _, set := range sets {
			cubes := bytes2.Split(set, []byte(","))
			for _, cube := range cubes {
				count, color := Split(cube, ' ')
				if bytes2.Equal(color, red) {
					r = max(r, count)
				} else if bytes2.Equal(color, green) {
					g = max(g, count)
				} else if bytes2.Equal(color, blue) {
					b = max(b, count)
				}
			}
		}
		valid_games += r * g * b
		k++
	}
	return valid_games
}

func Solve1Mad(bytes []byte) int {
	reader := bufio.NewReader(bytes2.NewBuffer(bytes))

	red := []byte("red")
	green := []byte("green")
	blue := []byte("blue")

	valid_games := 0
	k := 1
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			break
		}

		line = bytes2.Split(line, []byte(":"))[1]
		sets := bytes2.Split(line, []byte(";"))

		r, g, b := 0, 0, 0
		for _, set := range sets {
			cubes := bytes2.Split(set, []byte(","))
			for _, cube := range cubes {
				count, color := Split(cube, ' ')
				if bytes2.Equal(color, red) {
					r = max(r, count)
				} else if bytes2.Equal(color, green) {
					g = max(g, count)
				} else if bytes2.Equal(color, blue) {
					b = max(b, count)
				}
			}
		}
		if r <= 12 && g <= 13 && b <= 14 {
			valid_games += k
		}
		k++
	}
	return valid_games
}

func Split(input []byte, separator byte) (int, []byte) {
	start := 0
	for start < len(input) && unicode.IsSpace(rune(input[start])) {
		start++
	}

	end := len(input) - 1
	for end >= start && unicode.IsSpace(rune(input[end])) {
		end--
	}

	i := 0

	for start < len(input) && input[start] != separator {
		i = i*10 + int(input[start]-'0')
		start++
	}

	return i, input[start+1 : end+1]
}
