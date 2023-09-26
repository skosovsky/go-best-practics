package bench

import "testing"

func BenchmarkBench1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bench1()
	}
}

func BenchmarkBench2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bench2()
	}
}

func BenchmarkBench3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bench3()
	}
}

func BenchmarkBench4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bench4()
	}
}
