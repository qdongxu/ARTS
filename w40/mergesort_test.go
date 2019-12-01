package codekata

import "testing"

func TestMergeSort(t *testing.T) {
	verifySorter(MergeSort, true, t)
}

func BenchmarkMergesortShuffle(b *testing.B) {
	benchmarkShuffle(MergeSort, b)
}

func BenchmarkMergesortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(MergeSort, b)
}

func BenchmarkMergesortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(MergeSort, b)
}
