package main

import (
	"strconv"
	"strings"
	"sync"
)

type Range struct {
	start      int
	end        int
	startValue int
}

func (r Range) contains(val int) bool {
	return r.start <= val && val < r.end
}

func (r Range) value(val int) int {
	return r.startValue + (val - r.start)
}

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	seedLines := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int, len(seedLines))
	for i, c := range seedLines {
		seeds[i], _ = strconv.Atoi(c)
	}

	computedMap := make(map[int][]Range)
	k := -1

	for i, line := range lines {
		if i == 0 || line == "" {

			continue
		}
		if strings.Contains(line, "map") {
			k++
			//computedMap[k] = make(map[int]int)
			computedMap[k] = make([]Range, 0)
		} else {
			splt := strings.Split(line, " ")

			dest, _ := strconv.Atoi(splt[0])
			source, _ := strconv.Atoi(splt[1])
			rangeLen, _ := strconv.Atoi(splt[2])

			computedMap[k] = append(computedMap[k], Range{source, source + rangeLen, dest})
		}
	}

	currentMin := 0
	for i, seed := range seeds {
		val := seed
		for j := 0; j < k+1; j++ {
			for _, r := range computedMap[j] {
				if r.contains(val) {
					val = r.value(val)
					break
				}
			}
		}
		if i == 0 || val < currentMin {
			currentMin = val
		}
	}

	return currentMin
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	seedLines := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int, len(seedLines))
	for i, c := range seedLines {
		seeds[i], _ = strconv.Atoi(c)
	}

	computedMap := make(map[int][]Range)
	k := -1

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		if strings.Contains(line, "map") {
			k++
			computedMap[k] = make([]Range, 0)
		} else {
			splt := strings.Split(line, " ")

			dest, _ := strconv.Atoi(splt[0])
			source, _ := strconv.Atoi(splt[1])
			rangeLen, _ := strconv.Atoi(splt[2])

			computedMap[k] = append(computedMap[k], Range{source, source + rangeLen, dest})
		}
	}

	intervals := make([]Range, len(seeds)/2)
	for i := 0; i < len(seeds); i += 2 {
		intervals[i/2] = Range{seeds[i], seeds[i] + seeds[i+1], -1}
	}

	for i := 0; i < k+1; i++ {
		var toCheck = make([]Range, len(intervals))
		copy(toCheck, intervals)
		intervals = make([]Range, 0)
		for len(toCheck) > 0 {
			interval := toCheck[0]
			toCheck = toCheck[1:]
			found := false
			for _, r := range computedMap[i] {
				if r.contains(interval.start) && r.contains(interval.end) {
					intervals = append(intervals, Range{r.value(interval.start), r.value(interval.end), -1})
					found = true
					break
				} else if r.contains(interval.start) {
					intervals = append(intervals, Range{r.value(interval.start), r.value(r.end), -1})
					toCheck = append(toCheck, Range{r.end + 1, interval.end, -1})
					found = true
					break
				} else if r.contains(interval.end) {
					intervals = append(intervals, Range{r.value(r.start), r.value(interval.end), -1})
					toCheck = append(toCheck, Range{interval.start, r.start - 1, -1})
					found = true
					break
				} else if r.start < interval.start && interval.end < r.end {
					intervals = append(intervals, Range{r.value(r.start), r.value(r.end), -1})
					toCheck = append(toCheck, Range{interval.start, r.start - 1, -1})
					toCheck = append(toCheck, Range{r.end + 1, interval.end, -1})
					found = true
					break
				}
			}
			if !found {
				intervals = append(intervals, interval)
			}
		}
	}

	currentMin := int(^uint(0) >> 1)
	for _, interval := range intervals {
		currentMin = min(currentMin, interval.start)
	}
	return currentMin
}

func Solve2MT(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	lines := strings.Split(text, "\n")
	seedLines := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	seeds := make([]int, len(seedLines))
	for i, c := range seedLines {
		seeds[i], _ = strconv.Atoi(c)
	}

	computedMap := make(map[int][]Range)
	k := -1

	for i, line := range lines {
		if i == 0 || line == "" {
			continue
		}
		if strings.Contains(line, "map") {
			k++
			computedMap[k] = make([]Range, 0)
		} else {
			splt := strings.Split(line, " ")

			dest, _ := strconv.Atoi(splt[0])
			source, _ := strconv.Atoi(splt[1])
			rangeLen, _ := strconv.Atoi(splt[2])

			computedMap[k] = append(computedMap[k], Range{source, source + rangeLen, dest})
		}
	}

	var mu sync.Mutex
	currentMin := 9999999999999999

	var wg sync.WaitGroup

	for i := 0; i < len(seeds); i += 2 {
		from := seeds[i]
		to := seeds[i+1]

		wg.Add(1)
		go func(from, to int) {
			defer wg.Done()

			m := computeMin(computedMap, from, to, k)

			mu.Lock()
			defer mu.Unlock()

			if m < currentMin {
				currentMin = m
			}
		}(from, to)
	}

	wg.Wait()

	return currentMin
}

func computeMin(computedMap map[int][]Range, from, to, k int) int {
	currentMin := 0

	for j := from; j < from+to; j++ {
		seed := j
		val := seed
		for e := 0; e < k+1; e++ {
			for _, r := range computedMap[e] {
				if r.contains(val) {
					val = r.value(val)
					break
				}
			}
		}
		if (j == from) || val < currentMin {
			currentMin = val
		}
	}
	return currentMin
}
