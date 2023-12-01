package main

import (
	"testing"
)

func BenchmarkDay1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve1()
	}
}

func BenchmarkDay2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Solve2()
	}
}
