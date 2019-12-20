package w51

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestPriorityQueue(t *testing.T) {
	q := NewPriorityQueue(true)
	size := 10

	go func() {
		for i := size - 1; i >= 0; i-- {
			q.Push(&Data{Id: int32(i), Value: int32(i / 2)})
			time.Sleep(5 * time.Millisecond)
		}
	}()

	for i := size - 1; i >= 0; i-- {
		d := q.Take()
		//fmt.Printf("%v\n", d)
		if d.Id != int32(i) {
			t.Fatal(fmt.Sprintf("expect %d, but got: %d", i, d.Id))
		}
	}
}

func TestPriorityQueueReverse(t *testing.T) {
	q := NewPriorityQueue(true)
	size := 10000

	for i := 0; i < size; i++ {
		q.Push(&Data{Id: int32(i), Value: int32(i / 2)})
	}

	for i := size - 1; i >= 0; i-- {
		d := q.Take()
		if d.Id != int32(i) {
			t.Fatal(fmt.Sprintf("expect No.%d, but got %v", i, d.Id))
		}
	}
}

func TestPriorityQueueRandom(t *testing.T) {
	q := NewPriorityQueue(true)
	size := 100000

	for i := 0; i < size; i++ {
		q.Push(&Data{Id: int32(rand.Intn(size / 2)), Value: int32(i / 2)})
	}

	last := &Data{Id: int32(math.MaxInt32), Value: 0}
	for i := size - 1; i >= 0; i-- {
		d := q.Take()
		if d.Id > last.Id {
			t.Fatal(fmt.Sprintf("No.%d(%v) and No.%d(%v) are out of order", i-1, last, i, d))
		}
		last = d
	}
}

func BenchmarkPriorityQueue(b *testing.B) {
	q := NewPriorityQueue(true)
	size := 1000000
	for i := 0; i < size; i++ {
		id := rand.Intn(size / 2)
		q.Push(&Data{Id: int32(id), Value: i})
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		q.Take()
		id := rand.Intn(size / 2)
		q.Push(&Data{Id: int32(id), Value: i})
		//fmt.Printf("sum: %d\n", sum)
	}

}
