package main

import "fmt"

func main() {

	// var A = []int{5, 3, 2, 4, 5, 3, 2}

	var A []int
	for i := 0; i <= 499999; i++ {
		A = append(A, i+1)
	}

	A = append(A, A...)
	A = append(A, 5000000001)

	// fmt.Println(A)
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	var hash = map[int]int{}
	var result int

	for i := 0; i < len(A); i++ {
		if _, ok := hash[A[i]]; ok {
			// fmt.Println("Нашли ", A[i])
		} else {
			hash[A[i]] = 1
			result = A[i]
		}
	}
	return result
}
