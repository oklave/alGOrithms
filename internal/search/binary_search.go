package search

func BinarySearch(needle int, array []int) int {

	leftPointer := 0
	rightPointer := len(array) - 1

	for leftPointer <= rightPointer {
		median := (leftPointer + rightPointer) / 2

		if array[median] == needle {
			return needle
		} else if array[median] < needle {
			leftPointer = median + 1
		} else {
			rightPointer = median - 1
		}
	}

	return 0
}
