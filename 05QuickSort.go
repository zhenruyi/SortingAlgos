package main

import (
	"fmt"
)

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	QuickSort2(a, 0, len(a)-1)
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

func QuickSort2(a []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	key := a[left]
	for left < right {
		for right > left && a[right] >= key {
			right--
		}
		if right > left {
			a[left] = a[right]
			left++
		}
		for right > left && a[left] < key {
			left++
		}
		if right > left {
			a[right] = a[left]
			right--
		}
	}
	a[left] = key
	QuickSort2(a, i, left-1)
	QuickSort2(a, left+1, j)
}