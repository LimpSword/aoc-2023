package main

import (
	_ "embed"
	"fmt"
	"time"
)

//go:embed input.txt
var input []byte

func main() {
	t := time.Now()
	fmt.Println("1: ", Solve1(input))
	elapsed := time.Since(t)
	println("Solved Part 1 in", elapsed.Nanoseconds(), "ns")

	t = time.Now()
	fmt.Println("2: ", Solve2(input))
	elapsed = time.Since(t)
	println("Solved Part 2 in", elapsed.Nanoseconds(), "ns")
}
