package codekata

import (
	"math/rand"
	"testing"
	"time"
)

const benchSize = 32

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSelectSort(t *testing.T) {
	verifySorter(SelectSort, false, t)

}

func BenchmarkSelectsortShuffle(b *testing.B) {
	benchmarkShuffle(SelectSort, b)
}

func BenchmarkSelectsortNoShuffle(b *testing.B) {
	benchmarkNoShuffle(SelectSort, b)
}

func BenchmarkSelectsortNoShuffleAndReverse(b *testing.B) {
	benchmarkReverse(SelectSort, b)
}

func newArray(size int) []Data {
	array := make([]Data, size)
	for i := 0; i < size; i++ {
		array[i] = Data{Id: int32(i / 2), Value: int32(i)}
	}
	return array
}

func newShuffles(size int, N int) [][]Data {
	shuffles := make([][]Data, N)
	for i := 0; i < N; i++ {
		a := newArray(size)
		shuffle(a)
		shuffles[i] = a
	}
	return shuffles
}

func shuffle(array []Data) {
	size := len(array)
	for i := 0; i < size*2; i++ {
		m := rand.Intn(size)
		n := rand.Intn(size)
		if m == n {
			continue
		}

		array[m], array[n] = array[n], array[m]
	}

	for i := 0; i < size; i++ {
		array[i].Value = int32(i)
	}
}

func reverse(array []Data) {
	size := len(array)
	for left := 0; left < size/2; left++ {
		right := size - left - 1
		array[left], array[right] = array[right], array[left]
	}
}

func verifySorter(sorter func(array []Data), stable bool, t *testing.T) {
	sort := func(size int) {
		array := newArray(size)

		shuffle(array)

		// fmt.Printf("before: %v\n", array)
		sorter(array)
		// fmt.Printf("after: %v\n", array)

		for i, j := 0, 0; i < size-1; i++ {
			j++
			if array[i].Id > array[j].Id {
				t.Errorf("No. %d element(%v) is larger than No. %d element(%v)", i, array[i], j, array[j])
				break
			}

			if !stable {
				continue
			}

			if array[i].Id == array[j].Id {
				if array[i].Value.(int32) > array[j].Value.(int32) {
					t.Errorf("No. %d element(%v) is larger than No. %d element(%v)", i, array[i], j, array[j])
					break
				}
			}
		}
	}

	for _, i := range []int{1, 10, 1000} {
		// for _, i := range []int{10} {
		sort(i)
	}
}

func benchmarkShuffle(sorter func(array []Data), b *testing.B) {
	arrays := newShuffles(benchSize, b.N)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter(arrays[i])
	}

}

func benchmarkNoShuffle(sorter func(array []Data), b *testing.B) {
	array := newArray(benchSize)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		sorter(array)
	}

}

func benchmarkReverse(sorter func(array []Data), b *testing.B) {
	array := newArray(benchSize)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		reverse(array)
		sorter(array)
	}
}
