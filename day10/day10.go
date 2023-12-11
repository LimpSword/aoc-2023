package main

import (
	"strings"
)

func Solve1(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	graph := make([][]string, 0)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		graph = append(graph, strings.Split(line, ""))
	}

	start := []int{0, 0}
	for y, row := range graph {
		for x, cell := range row {
			if cell == "S" {
				start = []int{x, y}
			}
		}
	}

	visited := make(map[int]bool)
	bfs := func(graph [][]string, start []int) int {
		queue := make([][]int, 0)
		queue = append(queue, start)
		depth := 0
		for len(queue) > 0 {
			depth++
			next := make([][]int, 0)
			for _, node := range queue {
				x, y := node[0], node[1]
				visited[y*len(graph[0])+x] = true
				switch graph[y][x] {
				case "|":
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
				case "-":
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "L":
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "J":
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
				case "7":
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
				case "F":
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "S":
					// check only them (puzzle input)
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
				}
			}
			queue = next
		}
		return depth
	}
	return bfs(graph, start) - 1
}

func Solve2(bytes []byte) int {
	text := string(bytes)
	text = strings.ReplaceAll(text, "\r", "")

	graph := make([][]string, 0)
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		graph = append(graph, strings.Split(line, ""))
	}

	start := []int{0, 0}
	for y, row := range graph {
		for x, cell := range row {
			if cell == "S" {
				start = []int{x, y}
			}
		}
	}

	visited := make(map[int]bool)
	bfs := func(graph [][]string, start []int) int {
		queue := make([][]int, 0)
		queue = append(queue, start)
		depth := 0
		for len(queue) > 0 {
			depth++
			next := make([][]int, 0)
			for _, node := range queue {
				x, y := node[0], node[1]
				visited[y*len(graph[0])+x] = true
				switch graph[y][x] {
				case "|":
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
				case "-":
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "L":
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "J":
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
				case "7":
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
				case "F":
					if y < len(graph)-1 && !visited[(y+1)*len(graph[0])+x] {
						next = append(next, []int{x, y + 1})
					}
					if x < len(graph[y])-1 && !visited[(y)*len(graph[0])+x+1] {
						next = append(next, []int{x + 1, y})
					}
				case "S":
					// check only them (puzzle input)
					if x > 0 && !visited[(y)*len(graph[0])+x-1] {
						next = append(next, []int{x - 1, y})
					}
					if y > 0 && !visited[(y-1)*len(graph[0])+x] {
						next = append(next, []int{x, y - 1})
					}
				}
			}
			queue = next
		}
		return depth
	}

	bfs(graph, start)

	p2 := 0
	for y := range graph {
		for x := range graph[0] {
			if !visited[y*len(graph[0])+x] {
				count := 0
				for n := range graph[y][:x] {
					v := graph[y][n]
					// F and 7 represent the start and end of a loop
					if v == "F" || v == "7" || v == "|" {
						if visited[y*len(graph[0])+n] {
							count++
						}
					}
				}
				// inside if odd number of walls before
				if count%2 == 1 {
					p2 += 1
				}
			}
		}
	}
	return p2
}
