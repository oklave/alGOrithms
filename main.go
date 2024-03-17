package main

import (
	"fmt"

	"main.go/internal/search"
)

func main() {
	items := []int{1, 2, 4, 11, 31, 45, 63, 70, 100}
	fmt.Println(search.BinarySearch(63, items))
}
