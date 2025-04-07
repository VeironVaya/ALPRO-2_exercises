package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)
	pattern(n)
}
func pattern(n int) {
	var i int
	if n == 1 {
		fmt.Print("*")
	} else {
		pattern(n - 1)
		fmt.Println("")
		for i = 1; i <= n; i++ {
			fmt.Print("*")
		}

	}

}

// ALTERNATIVE
// func pattern(n int) {
// 	if n == 1 {
// 		printStars(1)
// 	} else {
// 		pattern(n - 1)
// 		fmt.Println()
// 		printStars(n)
// 	}
// }

// func printStars(n int) {
// 	if n == 0 {
// 		return
// 	}
// 	fmt.Print("*")
// 	printStars(n - 1)
// }
