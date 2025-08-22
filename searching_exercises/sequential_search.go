package main

type arrOfNum [10]int

func sequentialSearch(A arrOfNum, target int) int {
	var i int
	for i = 0; i < 10; i++ {
		if target == A[i] {
			return i
		}
	}
	return -1
}

func main(){
	var exArr arrOfNum

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

	sequentialSearch(exArr,11)
}
