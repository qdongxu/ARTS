package codekata

func HeapSort(array []Data) {
	len := len(array)

	for top := ((len - 1) - 1) / 2; top >= 0; top-- {
		heapify(array, top, len)
	}

	for top := len - 1; top >= 0; top-- {
		array[top], array[0] = array[0], array[top]
		len -= 1
		heapify(array, 0, len)
	}
}

func heapify(array []Data, top int, len int) {
	left := 2*top + 1
	right := 2*top + 2
	bigPos := top

	if left < len && array[left].Id > array[bigPos].Id {
		bigPos = left
	}
	if right < len && array[right].Id > array[bigPos].Id {
		bigPos = right
	}

	if bigPos != top {
		array[bigPos], array[top] = array[top], array[bigPos]
		heapify(array, bigPos, len)
	}
}
