package main

import "testing"

var res int

func benchmakrf1(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = f1(n)
	}

	res = r
}

func benchmakrf2(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = f2(n)
	}

	res = r
}

func benchmakrf3(b *testing.B, n int) {
	var r int
	for i := 0; i < b.N; i++ {
		r = f3(n)
	}

	res = r
}

func Benchmark30f1(b *testing.B) {
	benchmakrf1(b, 30)
}

func Benchmark10f1(b *testing.B) {
	benchmakrf1(b, 10)
}

func Benchmark30f2(b *testing.B) {
	benchmakrf2(b, 30)
}

func Benchmark10f2(b *testing.B) {
	benchmakrf2(b, 10)
}

func Benchmark30f3(b *testing.B) {
	benchmakrf3(b, 30)
}

func Benchmark10f3(b *testing.B) {
	benchmakrf3(b, 10)
}
