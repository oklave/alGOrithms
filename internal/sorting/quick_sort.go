package sorting

func QuickSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	pivot := array[0]
	var less, greater []int
	for _, num := range array[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	result := append(QuickSort(less), pivot)
	result = append(result, QuickSort(greater)...)
	return result
}
