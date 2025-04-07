package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fmt.Print(factorial(x), factorial(y), permute(x, y))

}

func factorial(num int) int {

	var result int = 1

	for num >= 1 {
		result = result * num
		num--
	}

	return result
}

func permute(x, y int) int {
	if x >= y {
		return factorial(x) / factorial(x-y)
	}
	return factorial(y) / factorial(y-x)
}
