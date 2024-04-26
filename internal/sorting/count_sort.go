package sorting

func getK(array []int) int {
	if len(array) == 0 {
		return 1
	}

	k := array[0]

	for _, v := range array {
		if v > k {
			k = v
		}
	}

	return k + 1
}

func CountingSort(array []int) {
	k := getK(array)
	array_of_counts := make([]int, k)

	for i := 0; i < len(array); i++ {
		array_of_counts[array[i]] += 1
	}

	for i, j := 0, 0; i < k; i++ {
		for {
			if array_of_counts[i] > 0 {
				array[j] = i
				j += 1
				array_of_counts[i] -= 1
				continue
			}
			break
		}
	}
}
