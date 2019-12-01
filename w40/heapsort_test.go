package codekata

import "testing"

func TestHeapSort(t *testing.T) {
	verifySorter(HeapSort, false, t)
}

func BenchmarkHeapsortShuffle(b *testing.B) {
	benchmarkShuffle(HeapSort, b)
}

func BenchmarkHeapsortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(HeapSort, b)
}

func BenchmarkHeapsortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(HeapSort, b)
}
