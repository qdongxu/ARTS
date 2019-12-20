package w49

import (
	"fmt"
	"sync"
)

const (
	initCapacity = 8
)

type (
	Data struct {
		Id    int32
		Value interface{}
	}

	CyclicArray struct {
		array    []*Data
		capacity int
		head     int
		tail     int
		size     int
		maxSize  int

		lock *sync.Mutex
	}
)

// NewCyclicArray creates and returns a new cyclic array.
//
// the length will be fixed to maxSize if maxSize is a non zero value. and Push will overwrite the earliest element
// when beyond one circle.
//
// the length will be extended automatically if maxSize is 0
//
// Note: maxSize should not be less than 0
//
func NewCyclicArray(maxSize int) *CyclicArray {
	if maxSize != 0 {
		return &CyclicArray{capacity: maxSize, maxSize: maxSize, array: make([]*Data, maxSize), lock: &sync.Mutex{}}
	}

	return &CyclicArray{capacity: initCapacity, maxSize: maxSize, array: make([]*Data, initCapacity), lock: &sync.Mutex{}}
}

func (c *CyclicArray) Get(index int) *Data {
	c.lock.Lock()
	defer c.lock.Unlock()

	realPos := c.index2RealPos(index)
	return c.array[realPos]
}

func (c *CyclicArray) index2RealPos(index int) int {
	if index >= c.size {
		panic(fmt.Sprintf("out of index: index: %d, size: %d", index, c.size))
	}
	index += c.head
	if index >= c.capacity {
		index -= c.capacity
	}
	return index
}

func (c *CyclicArray) Set(index int, data *Data) {
	c.lock.Lock()
	defer c.lock.Unlock()

	realPos := c.index2RealPos(index)

	c.array[realPos] = data
}

func (c *CyclicArray) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()

	return c.size
}

func (c *CyclicArray) Push(data *Data) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.isAutoExtend() && c.size >= c.capacity {
		c.extendSize()
	}

	if c.size < c.capacity {
		c.array[c.tail] = data
		c.tail++
		c.size++
		if c.tail >= c.capacity {
			c.tail = 0
		}

		return
	}

	// overwrite the head element as the fixed length array is full
	c.head++
	if c.head >= c.capacity {
		c.head = 0
	}

	c.array[c.tail] = data
	c.tail++
	if c.tail >= c.capacity {
		c.tail = 0
	}

}

func (c *CyclicArray) isAutoExtend() bool {
	return c.maxSize <= 0
}

func (c *CyclicArray) extendSize() {
	// fmt.Printf("resize: %d\n", c.capacity)
	newArray := make([]*Data, c.capacity*2)

	if c.tail > c.head {
		for i, j := 0, c.head; j < c.tail; i, j = i+1, j+1 {
			newArray[i] = c.array[j]
		}

	} else {
		i := 0
		for j := c.head; j < c.capacity; i, j = i+1, j+1 {
			newArray[i] = c.array[j]
		}

		for j := 0; j < c.tail; i, j = i+1, j+1 {
			newArray[i] = c.array[j]
		}
	}

	c.array = newArray
	c.capacity *= 2
	c.head = 0
	c.tail = c.size
}

func (c *CyclicArray) Take() (rel *Data, ok bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.size < 1 {
		return nil, false
	}

	rel = c.array[c.head]
	c.array[c.head] = nil

	c.size--
	c.head++
	if c.head >= c.capacity {
		c.head = 0
	}

	return rel, ok
}
