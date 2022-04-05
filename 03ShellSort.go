package main

import (
	"fmt"
)

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	ShellSort(a)
	fmt.Println(a)
}

func ShellSort(a []int) {
	length := len(a)
	interval := 1
	for interval < length / 3 {
		// 1 4 13 40
		interval = interval * 3 + 1
	}
	// 40 13 4 1
	for ; interval > 0; interval /= 3 {
		for i := interval; i < length; i++ {
			for j := i; j >= interval; j -= interval {
				if a[j - interval] > a[j] {
					a[j - interval], a[j] = a[j], a[j - interval]
				}
			}
		}
	}
}