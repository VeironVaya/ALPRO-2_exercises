package main

import (
	"fmt"
)

const numberOfTeams = 3

type ArrayWins [numberOfTeams]int

func inputDataWins(wins *ArrayWins, n *int) {
	var i int
	for i = 0; i < *n; i++ {
		fmt.Scan(&wins[i])
	}
}

func average(wins *ArrayWins, n int) float64 {
	var sum int
	var i int
	
	for i = 0; i < n; i++ {
		sum = sum + wins[i]
	}

	var result float64
	result = float64(sum) / float64(n)
	return result
}

func main() {
	var n int
	n = numberOfTeams

	var wins ArrayWins
	var avg float64

	inputDataWins(&wins, &n)
	avg = average(&wins, n)

	fmt.Print("\nAverage wins of ", n, " teams is ", avg)
}
