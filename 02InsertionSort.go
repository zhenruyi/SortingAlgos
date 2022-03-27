package main

import "fmt"

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	InsertionSort(a)
	fmt.Println(a)
}

func InsertionSort(a []int) {
	length := len(a)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			}
		}
	}
}
