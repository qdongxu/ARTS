package codekata

import "testing"

func TestFastSortv2(t *testing.T) {
	verifySorter(FastSort, false, t)
}

func BenchmarkFastsortv2Shuffle(b *testing.B) {
	benchmarkShuffle(FastSortV2, b)
}

func BenchmarkFastsortv2NoShuffle(b *testing.B) {
	benchmarkNoShuffle(FastSortV2, b)
}

func BenchmarkFastsortv2NoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(FastSortV2, b)
}
