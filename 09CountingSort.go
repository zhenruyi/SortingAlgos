package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := make([]int, 0)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		a = append(a, rand.Intn(100))
	}
	fmt.Println(a)
	CountingSort(a)
	fmt.Println(a)
}

func CountingSort(a []int) {
	count := map[int]int{}
	min, max := a[0], a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
		if val < min {
			min = val
		}
		count[val] += 1
	}
	index := 0
	for i := min; i <= max; i++ {
		for j := 0; j < count[i]; j++ {
			a[index] = i
			index++
		}
	}
}
