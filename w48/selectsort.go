package codekata

func SelectSort(array []Data) {
	for i := 0; i < len(array)-1; i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min].Id > array[j].Id {
				min = j
			}
		}

		if i != min {
			array[i], array[min] = array[min], array[i]
		}
	}
}

type Data struct {
	Id    int32
	Value interface{}
}
