package main

import "fmt"

type arrNum [10]int

func main() {
	var exArr arrNum
	var i, j, temp int

	exArr[0] = 9
	exArr[1] = 3
	exArr[2] = 2
	exArr[3] = 2
	exArr[4] = 1
	exArr[5] = 6
	exArr[6] = 8
	exArr[7] = 7
	exArr[8] = 11
	exArr[9] = 12

	fmt.Println(exArr)
	for i = 1; i < 10; i++ {
		temp = exArr[i]
		j = i - 1

		for j >= 0 && exArr[j] > temp {
			exArr[j+1] = exArr[j]
			j--
		}
		exArr[j+1] = temp

		fmt.Println("iteration:", i)
		fmt.Println(exArr)
		fmt.Println()

	}

}
