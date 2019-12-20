package codekata

import "testing"

func TestShellSort(t *testing.T) {
	verifySorter(ShellSort, false, t)
}

func BenchmarkShellsortShuffle(b *testing.B) {
	benchmarkShuffle(ShellSort, b)
}

func BenchmarkShellsortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(ShellSort, b)
}

func BenchmarkShellsortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(ShellSort, b)
}
