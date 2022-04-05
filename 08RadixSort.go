package main

import (
	"fmt"
)

func main() {
	a := []int{12, 3, 8, 5, 9, 11, 23, 36, 20, 28, 21}
	RadixSort2(a)
	fmt.Println(a)
}

func RadixSort(a []int) {
	key := maxOrderKey(a)
	tmp := make([]int, len(a))
	radix := 1
	for i := 0; i < key; i++ {
		count := [10]int{}
		for j := 0; j < len(a); j++ {
			value := (a[j] / radix) % 10
			count[value]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for j := len(a) - 1; j >= 0; j-- {
			value := (a[j] / radix) % 10
			tmp[count[value]-1] = a[j]
			count[value]--
		}
		for j := 0; j < len(a); j++ {
			a[j] = tmp[j]
		}
		radix *= 10
	}
}

// 错误的
func maxOrderKey(a []int) int {
	key := 1
	maxOrder := 10
	for _, value := range a {
		if value > maxOrder {
			maxOrder *= 10
			key++
		}
	}
	return key
}

func RadixSort2(a []int) {
	key := maxOrderKey2(a)
	radix := 1
	tmp := make([]int, len(a))
	for i := 0; i < key; i++ {
		count := [10]int{}
		for _, value := range a {
			value = (value / radix) % 10
			count[value]++
		}
		for j := 1; j < 10; j++ {
			count[j] += count[j-1]
		}
		for j := len(a)-1; j >= 0; j-- {
			value := (a[j] / radix) % 10
			tmp[count[value]-1] = a[j]
			count[value]--
		}
		for index, value := range tmp {
			a[index] = value
		}
		radix *= 10
	}
}

func maxOrderKey2(a []int) (key int) {
	maxValue := a[0]
	for _, value :=range a {
		if value > maxValue {
			maxValue = value
		}
	}
	key = 1
	radix := 10
	for radix < maxValue {
		key++
		radix *= 10
	}
	return
}