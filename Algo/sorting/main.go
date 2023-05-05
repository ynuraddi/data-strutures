package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	mass := []int{5, 2, 4, 6, 1, 3, 2, 6}
	for i := 0; i < 10000; i++ {
		mass = append(mass, rand.Int()%1000)
	}

	insertion := make([]int, len(mass))
	bubble := make([]int, len(mass))
	merge := make([]int, len(mass))

	copy(insertion, mass)
	copy(bubble, mass)
	copy(merge, bubble)

	tt := time.Now()
	insertionSort(insertion)
	insertionTime := time.Since(tt)
	fmt.Println("Insertion Sort: ", insertionTime)
	// fmt.Println(insertion)

	tt = time.Now()
	bubbleSort(bubble)
	bubbleTime := time.Since(tt)
	fmt.Println("Bubble Sort: ", bubbleTime)
	// fmt.Println(bubble)

	tt = time.Now()
	mergeSort(merge, 0, len(merge))
	mergeTime := time.Since(tt)
	fmt.Println("Merge Sort: ", mergeTime)
	// fmt.Println(merge)
}

func insertionSort(a []int) {
	for i := 1; i < len(a); i++ {
		k := a[i]
		j := i - 1
		for ; j >= 0 && a[j] > k; j-- {
			a[j+1] = a[j]
		}
		a[j+1] = k
	}
}

func bubbleSort(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func mergeSort(a []int, p, r int) {
	if p < r-1 {
		q := (p + r) >> 1
		mergeSort(a, p, q)
		mergeSort(a, q, r)
		merge(a, p, q, r)
	}
}

func merge(a []int, p, q, r int) {
	n1 := q - p
	n2 := r - q

	L := make([]int, n1)
	R := make([]int, n2)

	for i := 0; i < n1; i++ {
		L[i] = a[p+i]
	}
	for i := 0; i < n2; i++ {
		R[i] = a[q+i]
	}

	i, j := 0, 0
	for k := p; k < r; k++ {
		if i < len(L) && j < len(R) {
			if L[i] <= R[j] {
				a[k] = L[i]
				i++
			} else {
				a[k] = R[j]
				j++
			}
		} else {
			if i == len(L) {
				a[k] = R[j]
				j++
			} else {
				a[k] = L[i]
				i++
			}
		}
	}
}
