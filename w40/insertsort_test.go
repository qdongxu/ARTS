package codekata

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	verifySorter(InsertSort, true, t)
}

func BenchmarkInsertsortShuffle(b *testing.B) {
	benchmarkShuffle(InsertSort, b)
}

func BenchmarkInsertsortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(InsertSort, b)
}

func BenchmarkInsertsortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(InsertSort, b)
}
