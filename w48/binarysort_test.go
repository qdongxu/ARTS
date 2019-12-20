package codekata

import "testing"

func TestBinarySort(t *testing.T) {
	verifySorter(BinarySort, true, t)
}

func BenchmarkBinarysortShuffle(b *testing.B) {
	benchmarkShuffle(BinarySort, b)
}

func BenchmarkBinarysortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(BinarySort, b)
}

func BenchmarkBinarysortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(BinarySort, b)
}
