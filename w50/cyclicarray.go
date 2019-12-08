package codekata

import "fmt"

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
	}
)
func NewCyclicArray() *CyclicArray {
	return &CyclicArray{capacity: initCapacity, array: make([]*Data, initCapacity)}
}

func (c *CyclicArray) Get(index int) *Data {
	if index >= c.size {
		panic(fmt.Sprintf("out of index: index: %d, size: %d", index, c.size))
	}

	index += c.head
	if index >= c.capacity {
		index -= c.capacity
	}
	return c.array[index]
}

func (c *CyclicArray) Set(index int, data *Data) {
	index += c.head
	if index >= c.capacity {
		index -= c.capacity
	}

	c.array[index] = data
}

func (c *CyclicArray) Size() int {
	return c.size
}

func (c *CyclicArray) Push(data *Data) {
	if c.size >= c.capacity {
		fmt.Printf("resise: %d\n", c.capacity)
		newArray := make([]*Data, c.capacity*2)
		for i := 0; i < c.size; i++ {
			newArray[i] = c.Get(i)
		}

		c.array = newArray
		c.capacity *= 2
		c.head = 0
		c.tail = c.size

	}

	c.size++

	c.array[c.tail] = data
	c.tail++
	if c.tail >= c.capacity {
		c.tail = 0
	}
}

func (c *CyclicArray) Take() *Data {
	if c.size < 1 {
		return nil
	}

	rel := c.array[c.head]
	c.array[c.head] = nil

	c.size--
	c.head++
	if c.head >= c.capacity {
		c.head = 0
	}

	return rel
}
