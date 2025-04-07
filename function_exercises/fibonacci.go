package main

import "fmt"

func main() {
	var fibonacciVal, n int

	fmt.Scan(&n)
	fibonacciVal = fibonacci(n)
	fmt.Println(fibonacciVal)

}

func fibonacci(n int) int {

	var x, y int
	var i int
	var result, temp int

	x = 0
	y = 1

	if n == 1 {
		return x
	}

	result = x + y
	for i = 3; i <= n; i++ {
		temp = x + y
		x = y
		y = temp
		result = result + y
	}
	return result

}
