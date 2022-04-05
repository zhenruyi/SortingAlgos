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
	BucketSort(a)
	fmt.Println(a)
}

func BucketSort(a []int) {
	max := a[0]
	for _, val := range a {
		if val > max {
			max = val
		}
	}
	length := len(a)
	buckets := make([][]int, length)
	for _, val := range a {
		index := val * (length - 1) / max
		buckets[index] = append(buckets[index], val)
	}
	pos := 0
	for i := 0; i < length; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			sortInBucket(buckets[i])
			copy(a[pos:], buckets[i])
			pos += bucketLen
		}
	}
}

func sortInBucket(a []int) {
	length := len(a)
	for i := 0; i < length-1; i++ {
		min := i
		for j := i; j < length; j++ {
			if a[min] > a[j] {
				min = j
			}
		}
		a[min], a[i] = a[i], a[min]
	}
}