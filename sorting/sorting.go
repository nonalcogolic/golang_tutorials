package main

import (
	"fmt"
	"sort"
)

func main() {
	arrayA := [...]int{1, 2, 3, 4, 5, 6, 11}
	arrayB := [...]int{2, 5, 7, 9, 10}

	m := make(map[int]bool)

	for _, value := range arrayA {
		m[value] = false
	}

	for _, value := range arrayB {
		m[value] = false
	}

	slice_a := make([]int, 0)
	for key, _ := range m {
		slice_a = append(slice_a, key)
	}

	sort.Ints(slice_a)
	fmt.Println(slice_a)

}
