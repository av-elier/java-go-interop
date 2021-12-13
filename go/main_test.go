package main

import (
	"testing"

	"tradingview.com/java-go-interop/gosum"
	"tradingview.com/java-go-interop/javasum"
)

func BenchmarkSum(b *testing.B) {
	s := gosum.CreateGoSum()
	runSum(s, b)
}

func BenchmarkGraalSum(b *testing.B) {
	s := javasum.CreateJavaSum()
	runSum(s, b)
}
func BenchmarkSum10(b *testing.B) {
	s := gosum.CreateGoSum()
	runSumN(s, b, 10)
}
func BenchmarkGraalSum10(b *testing.B) {
	s := javasum.CreateJavaSum()
	runSumN(s, b, 10)
}

func BenchmarkSum1000(b *testing.B) {
	s := gosum.CreateGoSum()
	runSumN(s, b, 1000)
}

func BenchmarkGraalSum1000(b *testing.B) {
	s := javasum.CreateJavaSum()
	runSumN(s, b, 1000)
}

type summator interface {
	Calc(float64) float64
	CalcNTimes(int64, float64) float64
}

func runSum(s summator, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Calc(float64(i))
	}
}

func runSumN(s summator, b *testing.B, n int64) {
	for i := 0; i < b.N; i++ {
		s.CalcNTimes(n, float64(i))
	}
}
