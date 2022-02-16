package main

import "fmt"

func main() {

	var A = []int{3, 8, 9, 7, 6}
	var K = 1
	fmt.Println(Solution(A, K))
	// fmt.Println(A)
	// fmt.Println(K)
}

func Solution(A []int, K int) []int {
	var result = A
	for i := 0; i < K; i++ {
		result = OneRotation(A)
	}
	return result
}

func OneRotation(A []int) []int {
	var result = A
	// result[0] = A[len(A)-1]
	// fmt.Println(result[0])
	for i := 1; i < len(A); i++ {
		fmt.Println(A[i-1])
		result[i] = A[i-1]
		// fmt.Println(result[i])
	}
	return result
}
