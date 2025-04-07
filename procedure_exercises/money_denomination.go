package main

import "fmt"

func main() {
	var k10, k2, k1, money int
	fmt.Scan(&money)
	moneyDenom(money, &k10, &k2, &k1)
	fmt.Printf("%d 10000-bills\n%d 2000-bills\n%d 1000-bills", k10, k2, k1)

}

func moneyDenom(money int, k10, k2, k1 *int) {
	*k10 = money / 10000
	*k2 = money % 10000 / 2000
	*k1 = money % 10000 % 2000 / 1000
}
