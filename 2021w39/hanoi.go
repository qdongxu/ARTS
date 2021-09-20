package _021w39

import "fmt"

type Hanoi struct {
	name  string
	stack []int
}

func (h *Hanoi) Pop() int {
	if len(h.stack) > 0 {
		s := h.stack[0]
		h.stack = h.stack[1:]
		return s
	}
	panic(fmt.Errorf("pop from a empty stack"))
}

func (h *Hanoi) Push(s int) {
	h.stack = append([]int{s}, h.stack...)
}

func (h *Hanoi) Len() int {
	return len(h.stack)
}

func (h *Hanoi) String() string {
	return fmt.Sprintf("%s - %v", h.name, h.stack)

}

// Move all slices from a to c, hop on b
func Move(a *Hanoi, b *Hanoi, c *Hanoi, height int) {
	if height == 1 {
		s := a.Pop()
		c.Push(s)
	} else {
		Move(a, c, b, height-1)
		Move(a, b, c, 1)
		Move(b, a, c, height-1)
	}
	fmt.Printf("%v\n%v\n%v\n", a, b, c)
	fmt.Printf("\n\n")
}
