package main

import (
	_ "embed"
	"testing"
)

func BenchmarkPart1B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve1Mad(input)
		expected := 2006
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}

func BenchmarkPart2B(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2Mad(input)
		expected := 84911
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}
