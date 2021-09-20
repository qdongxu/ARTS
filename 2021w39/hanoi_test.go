package _021w39

import "testing"

var (
	A *Hanoi = &Hanoi{"A", []int{}}
	B *Hanoi = &Hanoi{"B", []int{}}
	C *Hanoi = &Hanoi{"C", []int{}}
)

func init() {
	for i := 7; i >= 0; i-- {
		A.Push(i)
	}
}
func TestMove(t *testing.T) {
	Move(A, B, C, A.Len())
	if A.Len() != 0 || B.Len() != 0 {
		t.Errorf("move failed\n")
	}

	for i := 0; i < 8; i++ {
		s := C.Pop()
		if s != i {
			t.Errorf("expected: %d, actual: %d", i, s)
		}
	}
}
