package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)
	fmt.Print(power(x, y))
}

func power(x, y int) int {
	if y == 1 {
		return x
	} else {
		return x * power(x, y-1)

	}

}
