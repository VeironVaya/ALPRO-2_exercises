package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(f(g(h(x))))
	fmt.Println(g(h(f(y))))
	fmt.Println(h(f(g(z))))
}

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}
