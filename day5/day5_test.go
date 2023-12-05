package main

import (
	_ "embed"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve1(input)
		expected := 227653707
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2(input)
		expected := 78775051
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}

func BenchmarkPart2MT(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2MT(input)
		expected := 78775051
		if r != expected {
			b.Errorf("Got %d, expected %d", r, expected)
		}
	}
}
