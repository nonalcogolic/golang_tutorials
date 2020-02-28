package main

import "fmt"

func main() {
	mass := [...]int{1, 2, 3, 4, 5, 6}
	slice_a := mass[:]
	fmt.Println(mass, slice_a)

	var minIdx int = 0
	var maxIdx int = 0
	var min int = 0
	var max int = 0
	if len(slice_a) > 0 {
		minIdx = 0
		maxIdx = 0
		min = slice_a[0]
		max = slice_a[0]
	}

	for i, value := range slice_a {
		if min > value {
			min = value
			minIdx = i
		}
		if max < value {
			max = value
			maxIdx = i
		}
	}

	if minIdx != maxIdx {
		temp := slice_a[minIdx]
		slice_a[minIdx] = slice_a[maxIdx]
		slice_a[maxIdx] = temp
	}

	fmt.Println(mass, slice_a)
}
