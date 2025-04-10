package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Print(fibonacci(n))
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)

}
