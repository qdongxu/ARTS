package w50

import (
	"math/rand"
	"testing"
)

func TestCyclicArray(t *testing.T) {
	array := NewCyclicArray(0)
	size := 3498
	for i := 0; i < size; i++ {
		array.Push(&Data{Id: int32(i), Value: int32(i)})
	}

	for i := 0; i < rand.Intn(1000)+7; i++ {
		v, _ := array.Take()
		array.Push(v)
	}

	for i := 0; i < size-1; i++ {
		curr := *array.Get(i)
		next := *array.Get(i + 1)
		if next.Id == curr.Id+1 {
			// OK
		} else if next.Id == 0 {
			// OK
		} else {
			t.Fatalf("elements out of order: No. %d(%v), No. %d(%v) ", i, curr, i+1, next)
		}
	}
}

func TestCyclicArrayFixLength(t *testing.T) {
	maxSize := 10
	array := NewCyclicArray(10)
	size := 3498
	for i := 0; i < size; i++ {
		array.Push(&Data{Id: int32(i), Value: int32(i)})
	}

	small := size - maxSize

	for i := 0; i < rand.Intn(100)+7; i++ {
		v, _ := array.Take()
		array.Push(v)
	}

	for i := 0; i < maxSize-1; i++ {
		curr := *array.Get(i)
		next := *array.Get(i + 1)
		if next.Id == curr.Id+1 {
			// OK
		} else if next.Id == int32(small) {
			// OK
		} else {
			t.Fatalf("elements out of order: No. %d(%v), No. %d(%v) ", i, curr, i+1, next)
		}
	}
}

func BenchmarkCyclicArray(b *testing.B) {
	array := NewCyclicArray(0)
	for i := 0; i < 17; i++ {
		array.Push(&Data{Id: int32(i), Value: int32(i)})
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < rand.Intn(100)+7; i++ {
			v, _ := array.Take()
			array.Push(v)
		}
	}

}
