package main

import "fmt"

type ArrNum [4]int

func main() {
	var A ArrNum
	var n int

	fmt.Scan(&n)
	inpArr(&A, n)
	printArray(A)
	reverseArray(&A)
	printArray(A)

	fmt.Println("option one palindrom function:")
	if isPalindrom(A) {
		fmt.Println("Array is palindrom")
	} else {
		fmt.Println("Array isn't palindrom")
	}

	fmt.Println("option two palindrom function:")
	if isPalindromAlternative(A) {
		fmt.Println("Array is palindrom")
	} else {
		fmt.Println("Array isn't palindrom")
	}

}

func inpArr(A *ArrNum, n int) {
	var i int

	for i = 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
}

func reverseArray(A *ArrNum) {
	var i int = 0
	var temp int
	var end int = len(A) - 1

	for i < end {
		temp = A[i]
		A[i] = A[end]
		A[end] = temp
		i++
		end--
	}

}

func printArray(A ArrNum) {
	var i int
	fmt.Println("Array anda")
	for i = 0; i < len(A); i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println("")

}

func isPalindrom(A ArrNum) bool {
	var i int = 0

	var end int = len(A) - 1

	for i < end {
		if A[i] != A[end] {
			return false
		}
		i++
		end--

	}
	return true
}

func isPalindromAlternative(A ArrNum) bool {
	var temp ArrNum
	temp = A
	reverseArray(&temp)
	return A == temp

}
