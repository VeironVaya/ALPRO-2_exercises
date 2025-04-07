package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	oddNumSer(n)
}

func oddNumSer(n int) {
	if n == 1 {
		fmt.Print(1)
	} else {
		oddNumSer(n - 1)
		if n%2 != 0 {
			fmt.Print(" ", n)
		}

	}

}
