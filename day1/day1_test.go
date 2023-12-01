package main

import (
	"testing"
)

func BenchmarkDay1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve1()
		if r != 55447 {
			b.Errorf("Got %d, expected 55447", r)
		}
	}
}

func BenchmarkDay2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2()
		if r != 54706 {
			b.Errorf("Got %d, expected 54706", r)
		}
	}
}

func BenchmarkDay2Mad(b *testing.B) {
	for n := 0; n < b.N; n++ {
		r := Solve2Mad()
		if r != 54706 {
			b.Errorf("Got %d, expected 54706", r)
		}
	}
}
