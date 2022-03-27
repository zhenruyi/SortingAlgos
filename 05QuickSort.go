package main

import (
	"fmt"
)

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	QuickSort(a, 0, len(a)-1)
	fmt.Println(a)
}

func QuickSort(a []int, left int, right int) {
	if left >= right {
		return
	}

	i, j, key := left, right, a[left]
	for i < j {
		for i < j && key < a[j] {
			j--
		}
		if i < j {
			a[i] = a[j]
			i++
		}

		for i < j && key > a[i] {
			i++
		}
		if i < j {
			a[j] = a[i]
			j--
		}
	}
	a[i] = key
	QuickSort(a, left, i-1)
	QuickSort(a, i+1, right)
}
