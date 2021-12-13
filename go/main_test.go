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

type summator interface {
	Calc(float64) float64
}

func runSum(s summator, b *testing.B) {
	for i := 0; i < b.N; i++ {
		s.Calc(float64(i))
	}
}
