package main

import (
	_ "embed"
	"testing"
)

func BenchmarkPart1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve1(input)
		if r != 55447 {
			b.Errorf("Got %d, expected 55447", r)
		}
	}
}

func BenchmarkPart2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2Mad(input)
		if r != 54706 {
			b.Errorf("Got %d, expected 54706", r)
		}
	}
}
