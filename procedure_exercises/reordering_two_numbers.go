package main

import "fmt"

func main() {
	var a, b int
	var stop bool = false
	for !stop {
		fmt.Scan(&a, &b)
		reordering(&a, &b)
		fmt.Println(a, b)
		stop = a == b
	}

}

func reordering(a, b *int) {
	var temp int
	temp = *b
	if *a > *b {
		*b = *a
		*a = temp
	}

}
