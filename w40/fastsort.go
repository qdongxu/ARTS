package codekata

func FastSort(array []Data) {
	size := len(array)
	dofastsort(array, 0, size-1)
}

func dofastsort(array []Data, start int, end int) {
	mid := start + (end-start)/2
	midvalue := array[mid]

	lo := start
	hi := end

	for lo < hi {
		for lo <= end && array[lo].Id < midvalue.Id {
			lo++
		}

		for hi >= start && array[hi].Id > midvalue.Id {
			hi--
		}

		if lo < hi {
			array[lo], array[hi] = array[hi], array[lo]
			lo++
			hi--
		} else if lo == hi {
			lo++
			hi--
		}

	}

	if hi > start {
		dofastsort(array, start, hi)
	}

	if lo < end {
		dofastsort(array, lo, end)
	}
}
