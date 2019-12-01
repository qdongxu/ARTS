package codekata

func ShellSort(array []Data) {
	for _, step := range steps(len(array)) {
		size := len(array)
		for i := 0 + step; i < size; i++ {
			tmp := array[i]
			for j := i; j > 0 && j-step >= 0; j -= step {
				if tmp.Id >= array[j-step].Id {
					break
				}

				array[j] = array[j-step]
				array[j-step] = tmp
			}
		}
	}
}

func steps(n int) []int {
	results := make([]int, 0)
	for i := 1; i < n; i = (i+1)*2 - 1 {
		results = append(results, i)
	}

	size := len(results)
	for left, right := 0, size-1; left < right; {
		results[left], results[right] = results[right], results[left]

		left++
		right--
	}

	return results
}
