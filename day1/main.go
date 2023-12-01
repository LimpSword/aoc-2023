package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("1: ", Solve1())
	elapsed := time.Since(t)
	println("Solved Part 1 in", elapsed.Nanoseconds(), "ns")

	t = time.Now()
	fmt.Println("2: ", Solve2())
	elapsed = time.Since(t)
	println("Solved Part 2 in", elapsed.Nanoseconds(), "ns")
}
