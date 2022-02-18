package main

import "fmt"

func main() {

	var A = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 9}

	// var A []int
	// for i := 0; i <= 100000; i++ {
	// 	if i != 54356 {
	// 		A = append(A, i+1)
	// 	}
	// }
	fmt.Println(A)
	fmt.Println(Solution(A))
}

func Solution(A []int) int {

	if doubleSeek(A) == 1 {
		return 0
	}

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
	fmt.Println("max-min ", max-min+1)
	fmt.Println("len ", len(A))
	if max-min+1 == len(A) {
		return 1
	}
	return 0
}

func doubleSeek(A []int) int {
	var hash = map[int]int{}

	for i := 0; i < len(A); i++ {
		if _, ok := hash[A[i]]; ok {
			return 1
		} else {
			hash[A[i]] = 1
		}
	}
	return 0
}
