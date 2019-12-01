package codekata

func BinarySort(array []Data) {
	dobinarysort(array, 0, len(array)-1)
}

func dobinarysort(array []Data, start int, end int) {
	for i := start + 1; i <= end; i++ {
		tmp := array[i]
		pos := bianrySearch(array, start, i-1, tmp.Id)
		j := i
		for ; j > pos; j-- {
			array[j] = array[j-1]
		}

		if pos != i {
			array[j] = tmp
		}
	}
}

// binarySearch returns a index, the new item should be inserted before this index
// if lastHi equals end, that means no insert needed
func bianrySearch(array []Data, start int, end int, id int32) (lastHi int) {
	lastHi = end + 1
	for start <= end {
		mid := start + (end-start)/2
		if array[mid].Id <= id {
			start = mid + 1
		} else {
			lastHi = mid
			end = mid - 1
		}
	}

	return lastHi

}
