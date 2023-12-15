package main

import (
	"slices"
	"strconv"
	"strings"
)

func hash(text string) int {
	result := 0
	for _, c := range text {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	all := strings.Split(text, ",")

	result := 0
	for i := 0; i < len(all); i++ {
		result += hash(all[i])
	}

	return result
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\n", "")

	all := strings.Split(text, ",")
	hashmap := make(map[int]map[string]int)
	orderMap := make(map[int][]string)

	for i := 0; i < len(all); i++ {
		if strings.ContainsRune(all[i], '=') {
			label := strings.Split(all[i], "=")[0]
			focal := strings.Split(all[i], "=")[1]
			box := hash(label)
			if hashmap[box] == nil {
				hashmap[box] = make(map[string]int)

				orderMap[box] = make([]string, 0)
				orderMap[box] = append(orderMap[box], label)
			}
			focalInt, _ := strconv.Atoi(focal)
			hashmap[box][label] = focalInt

			if !slices.Contains(orderMap[box], label) {
				orderMap[box] = append(orderMap[box], label)
			}
		} else {
			label := strings.Split(all[i], "-")[0]
			box := hash(label)
			if hashmap[box] != nil {
				// remove label from order map
				for j := 0; j < len(orderMap[box]); j++ {
					if orderMap[box][j] == label {
						orderMap[box] = append(orderMap[box][:j], orderMap[box][j+1:]...)
						break
					}
				}
				delete(hashmap[box], label)
			}
		}
	}

	power := 0
	for i := range hashmap {
		for j := 0; j < len(orderMap[i]); j++ {
			power += (i + 1) * (j + 1) * hashmap[i][orderMap[i][j]]
		}
	}

	return power
}
