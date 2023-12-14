package main

import (
	_ "embed"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve1(input)
		expected := 30575
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2(input)
		expected := 37478
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}
