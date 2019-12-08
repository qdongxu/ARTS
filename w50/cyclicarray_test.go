package codekata

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestCyclicArray(t *testing.T) {
	array := NewCyclicArray()
	for i := 0; i < 17; i++ {
		array.Push(&Data{Id: int32(i), Value: int32(i)})
	}

	for i := 0; i < rand.Intn(100)+7; i++ {
		v := array.Take()
		array.Push(v)
	}

	for i := 0; i < 17; i++ {
		fmt.Printf("%v, ", *array.Get(i))
	}

}

func BenchmarkCyclicArray(b *testing.B) {
	array := NewCyclicArray()
	for i := 0; i < 17; i++ {
		array.Push(&Data{Id: int32(i), Value: int32(i)})
	}


	for i := 0; i < b.N; i++ {
		for i := 0; i < rand.Intn(100)+7; i++ {
			v := array.Take()
			array.Push(v)
		}
	}

}
