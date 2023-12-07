package main

import (
	"AoC-2023/day1"
	"AoC-2023/day2"
	"AoC-2023/day3"
	"AoC-2023/day4"
	"AoC-2023/day5"
	"AoC-2023/day6"
	"AoC-2023/day7"
	"testing"
)

func BenchmarkDay1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.Solve(true)
	}
}

func BenchmarkDay2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day2.Solve(true)
	}
}

func BenchmarkDay3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day3.Solve(true)
	}
}

func BenchmarkDay4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day4.Solve(true)
	}
}

func BenchmarkDay5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day5.Solve(true)
	}
}

func BenchmarkDay6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day6.Solve(true)
	}
}

func BenchmarkDay7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day7.Solve(true)
	}
}
