package codekata

func BubbleSort(array []Data) {
	size := len(array)
	for i := 0; i < size-1; i++ {
		for j := 0; j < size-1-i; j++ {
			if array[j].Id > array[j+1].Id {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}
