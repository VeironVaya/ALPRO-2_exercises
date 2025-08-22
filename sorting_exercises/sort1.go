package main

import "fmt"

type arrNum [10]int

func main() {
	var exArr arrNum
	var i, j, minIdx, temp int

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
	for i = 0; i < 9; i++ {
		minIdx = i
		for j = i + 1; j < 10; j++ {
			if exArr[j] < exArr[minIdx] {
				minIdx = j
			}
		}
		
		temp = exArr[i]
		exArr[i] = exArr[minIdx]
		exArr[minIdx] = temp

		fmt.Println("iteration:", i+1)
		fmt.Println(exArr)
		fmt.Println()
	}
}
