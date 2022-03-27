package main

import "fmt"

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	SelectionSort(a)
	fmt.Println(a)
}

func SelectionSort(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i+1; j < length; j++ {
			if a[j] < a[min] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
}
