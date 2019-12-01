package codekata

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	verifySorter(BubbleSort, true, t)
}

func BenchmarkBubblesortShuffle(b *testing.B) {
	benchmarkShuffle(BubbleSort, b)
}

func BenchmarkBubblesortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(BubbleSort, b)
}

func BenchmarkBubblesortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(BubbleSort, b)
}
