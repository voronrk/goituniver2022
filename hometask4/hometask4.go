package main

import "fmt"

func main() {

	// var A = []int{15, 13, 12, 14, 16, 17, 9}

	var A []int
	for i := 0; i <= 100000; i++ {
		if i != 54356 {
			A = append(A, i+1)
		}
	}
	// fmt.Println(A)
	fmt.Println(Solution(A))
}

func Solution(A []int) int {

	min := 1000000000
	max := 1

	for i := 0; i < len(A); i++ {
		if min > A[i] {
			min = A[i]
		}
		if max < A[i] {
			max = A[i]
		}
	}
	fmt.Println(min)
	fmt.Println(max)
	if max-min+1 == len(A) {
		return 1
	}
	return 0
}
