package codekata

import "testing"

func TestFastSort(t *testing.T) {
	verifySorter(FastSort, false, t)
}

func BenchmarkFastsortShuffle(b *testing.B) {
	benchmarkShuffle(FastSort, b)
}

func BenchmarkFastsortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(FastSort, b)
}

func BenchmarkFastsortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(FastSort, b)
}
