package main

import "fmt"

func main() {

	// var A = []int{5, 3, 2, 4, 1, 7}
	var A []int

	for i := 0; i <= 100000; i++ {
		if i != 99999 {
			A = append(A, i+1)
		}
	}

	fmt.Println(len(A))
	fmt.Println(Solution(A))
}

func Solution(A []int) int {
	for i := 1; i < len(A)+1; i++ {
		if check(A, i) {
			return i
		}
	}
	return 0
}

func check(A []int, num int) bool {
	for i := 0; i < len(A); i++ {
		if num == A[i] {
			return false
		}
	}
	return true
}
