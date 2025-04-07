package main

import "fmt"

func main() {
	var a, b, c, d int
	var permuteVal1, permuteVal2, combinationVal1, combinationVal2 int
	fmt.Scan(&a, &b, &c, &d)
	permuteVal1 = permute(a, c)
	combinationVal1 = combination(a, c)
	permuteVal2 = permute(b, d)
	combinationVal2 = combination(b, d)

	fmt.Println(permuteVal1, combinationVal1)
	fmt.Println(permuteVal2, combinationVal2)

}

func factorial(num int) int {

	var result int = 1

	for num >= 1 {
		result = result * num
		num--
	}

	return result
}

func permute(x, y int) int {
	return factorial(x) / factorial(x-y)
}

func combination(x, y int) int {
	return factorial(x) / (factorial(y) * factorial(x-y))
}
