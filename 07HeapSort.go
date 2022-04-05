package main

import "fmt"

func main() {
	a := []int{2, 6, 3, 6, 7, 29, 22, 15}
	HeapSort(a)
	fmt.Println(a)
}

func HeapSort(a []int) {
	// 构建最大堆
	for i := len(a)/2 - 1; i >= 0; i-- {
		AdjustHeap(a, i, len(a))
	}
	// 置于队尾
	for j := len(a) - 1; j > 0; j-- {
		a[0], a[j] = a[j], a[0]
		AdjustHeap(a, 0, j)
	}
}

func AdjustHeap(a []int, i, length int) {
	tmp := a[i]
	for k := i*2 + 1; k < length; k = k*2 + 1 {
		if k+1 < length && a[k] < a[k+1] {
			k++
		}
		if a[k] > tmp {
			a[i] = a[k]
			i = k
		} else {
			break
		}
	}
	a[i] = tmp
}
