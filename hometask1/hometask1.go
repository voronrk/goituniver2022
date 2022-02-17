package main

import "fmt"

func main() {

	var A = []int{3, 8, 9, 7, 6}
	fmt.Println("i", A)
	for i := 0; i <= 10; i++ {
		fmt.Println(i, Solution(A, i))
	}
}

func Solution(A []int, K int) []int {
	var newIndex int
	var result []int

	K = K % len(A)

	for i := len(A) - K; i < len(A)-K+len(A); i++ {
		newIndex = i - ((i)/len(A))*len(A)
		result = append(result, A[newIndex])
	}
	return result
}
