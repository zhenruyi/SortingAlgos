package main

import "fmt"

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	res := MergeSort(a)
	fmt.Println(res)
}

func MergeSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	middle := len(a) / 2
	left := a[:middle]
	right := a[middle:]
	return Merge(MergeSort(left), MergeSort(right))
}

func Merge(left []int, right []int) []int {
	var result []int
	for len(left) > 0 && len(right) > 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		} else {
			result = append(result, left[0])
			left = left[1:]
		}
	}
	for len(left) > 0 {
		result = append(result, left[0])
		left = left[1:]
	}
	for len(right) > 0 {
		result = append(result, right[0])
		right = right[1:]
	}
	return result
}
