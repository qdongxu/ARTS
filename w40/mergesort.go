package codekata

func MergeSort(array []Data) {
	size := len(array)
	buf := make([]Data, size)
	domergesort(array, 0, size-1, buf)
}

func domergesort(array []Data, left int, right int, buf []Data) {
	if left < right {
		mid := left + (right-left)/2
		domergesort(array, left, mid, buf)
		domergesort(array, mid+1, right, buf)
		merge(array, left, mid, right, buf)
	}
}

func merge(array []Data, left int, mid int, right int, buf []Data) {
	l, r := left, mid+1
	i := 0
	for l <= mid && r <= right {
		if array[l].Id <= array[r].Id {
			buf[i] = array[l]
			i++
			l++
		} else {
			buf[i] = array[r]
			i++
			r++
		}
	}

	for l <= mid {
		buf[i] = array[l]
		i++
		l++
	}

	for r <= right {
		buf[i] = array[r]
		i++
		r++
	}

	j := 0
	for i := left; i <= right; i++ {
		array[i] = buf[j]
		j++
	}
}
