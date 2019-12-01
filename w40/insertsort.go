package codekata

func InsertSort(array []Data) {
	doinsertsort(array, 0, len(array)-1)
}

func doinsertsort(array []Data, start int, end int) {
	for i := start + 1; i <= end; i++ {
		tmp := array[i]
		for j := i; j > 0; j-- {
			if tmp.Id >= array[j-1].Id {
				break
			}

			array[j] = array[j-1]
			array[j-1] = tmp
		}
	}
}
