package sorting

func MergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}
	mid := len(array) / 2
	left := MergeSort(array[:mid])
	right := MergeSort(array[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	var merged []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] > 0 && len(right) > 0 {
			merged = append(merged, left[0])
			left = left[1:]
		} else {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}
	merged = append(merged, left...)
	merged = append(merged, right...)
	return merged
}
