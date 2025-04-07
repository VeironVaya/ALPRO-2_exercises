package main

import "fmt"

func main() {
	var n, temp int

	fmt.Scan(&n)
	temp = n
	factor(n, temp)
}

func factor(n, temp int) {
	if n == 1 {
		fmt.Print(1)
	} else {
		factor(n-1, temp)
		if temp%n == 0 {
			fmt.Print(" ", n)
		}

	}

}
