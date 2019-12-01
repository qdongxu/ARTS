package codekata

func FastSortV2(array []Data) {
	size := len(array)
	dofastsortv2(array, 0, size-1)
}

func dofastsortv2(array []Data, start int, end int) {

	if end-start <= 64 {
		doinsertsort(array, start, end)
		return
	}

	lo := start
	hi := end
	mid := start + (end-start)/2
	midvalue := array[mid]

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
		dofastsortv2(array, start, hi)
	}

	if lo < end {
		dofastsortv2(array, lo, end)
	}
}
