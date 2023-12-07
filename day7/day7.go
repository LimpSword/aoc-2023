package main

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
)

var cardByScore map[int]string

func init() {
	cardByScore = make(map[int]string)
}

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	hands := make([]string, len(lines))
	bids := make(map[string]int)
	for k, line := range lines {
		values := strings.Split(line, " ")
		bid, _ := strconv.Atoi(values[1])
		hands[k] = values[0]
		bids[values[0]] = bid
	}

	sortedHands := orderFromType(hands)
	finalList := make([]string, 0)
	for _, hand := range sortedHands {
		finalList = append(finalList, cardByScore[hand])
	}

	result := 0
	for i := 0; i < len(finalList); i++ {
		result += bids[finalList[i]] * (i + 1)
	}

	return result
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	hands := make([]string, len(lines))
	bids := make(map[string]int)
	for k, line := range lines {
		values := strings.Split(line, " ")
		bid, _ := strconv.Atoi(values[1])
		hands[k] = values[0]
		bids[values[0]] = bid
	}

	sortedHands := orderFromType2(hands)
	finalList := make([]string, 0)
	for _, hand := range sortedHands {
		finalList = append(finalList, cardByScore[hand])
	}

	result := 0
	for i := 0; i < len(finalList); i++ {
		result += bids[finalList[i]] * (i + 1)
	}

	return result
}

func orderFromType(hands []string) []int {
	cardByScore = make(map[int]string)
	sortedHands := make([]int, len(hands))
	for k, hand := range hands {
		score := getScore(hand)
		sortedHands[k] = score
		cardByScore[score] = hand
	}
	slices.Sort(sortedHands)
	return sortedHands
}

func orderFromType2(hands []string) []int {
	cardByScore = make(map[int]string)
	sortedHands := make([]int, len(hands))
	for k, hand := range hands {
		score := getScore2(hand)
		sortedHands[k] = score
		cardByScore[score] = hand
	}
	slices.Sort(sortedHands)
	return sortedHands
}

func getScore(hand string) int {
	score := 0
	cardSet := Unique(strings.Split(hand, ""))
	if len(cardSet) == 1 {
		score = 7 // five of a kind
	} else if strings.Count(hand, cardSet[0]) == 4 || strings.Count(hand, cardSet[1]) == 4 {
		score = 6 // four of a kind
	} else if len(cardSet) == 2 {
		score = 5 // full house
	} else if strings.Count(hand, cardSet[0]) == 3 || strings.Count(hand, cardSet[1]) == 3 || strings.Count(hand, cardSet[2]) == 3 {
		score = 3 // three of a kind
	} else if (strings.Count(hand, cardSet[0]) == 2 && strings.Count(hand, cardSet[1]) == 2) || (strings.Count(hand, cardSet[0]) == 2 && strings.Count(hand, cardSet[2]) == 2) || (strings.Count(hand, cardSet[1]) == 2 && strings.Count(hand, cardSet[2]) == 2) {
		score = 2 // two pairs
	} else if strings.Count(hand, cardSet[0]) == 2 || strings.Count(hand, cardSet[1]) == 2 || strings.Count(hand, cardSet[2]) == 2 || strings.Count(hand, cardSet[3]) == 2 {
		score = 1 // one pair
	}

	score2 := ""
	for _, c := range hand {
		if c == 'A' {
			score2 += "99"
			continue
		} else if c == 'K' {
			score2 += "98"
			continue
		} else if c == 'Q' {
			score2 += "97"
			continue
		} else if c == 'J' {
			score2 += "96"
			continue
		} else if c == 'T' {
			score2 += "95"
			continue
		} else if unicode.IsDigit(c) {
			v := int(c - '0')
			score2 += strconv.Itoa(v * 10)
			continue
		}
	}
	scoreInt, _ := strconv.Atoi(strconv.Itoa(score) + score2)

	return scoreInt
}

func getScore2(hand string) int {
	score := 0
	cardSet := Unique(strings.Split(hand, ""))
	h2 := strings.ReplaceAll(hand, "", "")

	jokerCount := strings.Count(hand, "J")
	if jokerCount == 5 {
		hand = strings.ReplaceAll(hand, "J", "A")
	} else {
		mostCommon := ""
		for _, c := range cardSet {
			if strings.Count(hand, c) > strings.Count(hand, mostCommon) {
				mostCommon = c
			}
		}
		if mostCommon == "J" {
			h := strings.ReplaceAll(hand, "J", "")
			mostCommon := ""
			for _, c := range cardSet {
				if strings.Count(h, c) > strings.Count(h, mostCommon) {
					mostCommon = c
				}
			}
			hand = strings.ReplaceAll(hand, "J", mostCommon)
		}
	}
	for i := 0; i < len(cardSet); i++ {
		if cardSet[i] == "J" {
			cardSet = append(cardSet[:i], cardSet[i+1:]...)
			i--
		}
	}

	if len(cardSet) == 0 {
		score = 7 // five of a kind
	} else {
		h := strings.ReplaceAll(hand, "J", "")
		mostCommon := cardSet[0]
		for _, c := range cardSet {
			if strings.Count(h, c) > strings.Count(h, mostCommon) {
				mostCommon = c
			}
		}

		if len(cardSet) == 1 || jokerCount == 4 {
			score = 7 // five of a kind
		} else if strings.Count(hand, mostCommon)+jokerCount == 4 {
			score = 6 // four of a kind
		} else if len(cardSet) == 2 {
			score = 5 // full house
		} else if strings.Count(hand, mostCommon)+jokerCount == 3 {
			score = 3 // three of a kind
		} else if len(cardSet) == 3 && strings.Count(hand, mostCommon)+jokerCount == 2 {
			// surely not possible with jokers
			score = 2 // two pairs
		} else if len(cardSet) == 4 && strings.Count(hand, mostCommon)+jokerCount == 2 {
			score = 1 // one pair
		}
	}

	score2 := ""
	for _, c := range h2 {
		if c == 'A' {
			score2 += "99"
			continue
		} else if c == 'K' {
			score2 += "98"
			continue
		} else if c == 'Q' {
			score2 += "97"
			continue
		} else if c == 'T' {
			score2 += "95"
			continue
		} else if c == 'J' {
			score2 += "10"
			continue
		} else if unicode.IsDigit(c) {
			v := int(c - '0')
			score2 += strconv.Itoa((v * 10) + 1)
			continue
		}
	}
	scoreInt, _ := strconv.Atoi(strconv.Itoa(score) + score2)

	return scoreInt
}

func Unique(slice []string) []string {
	uniqMap := make(map[string]struct{})
	for _, v := range slice {
		uniqMap[v] = struct{}{}
	}

	uniqSlice := make([]string, 0, len(uniqMap))
	for v := range uniqMap {
		uniqSlice = append(uniqSlice, v)
	}
	return uniqSlice
}
