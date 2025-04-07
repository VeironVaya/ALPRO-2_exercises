package main

import "fmt"

func main() {
	var purchase, discountVal, paidVal int
	var member bool

	fmt.Scan(&purchase, &member)
	discountVal = discount(purchase, member)
	paidVal = paid(purchase, discountVal)
	fmt.Print(paidVal)

}

func discount(p int, m bool) int {
	if p > 100000 && m == true {
		return 10
	} else if p > 100000 && m == false {
		return 5
	}
	return 0
}

func paid(p int, disc int) int {
	if disc != 0 {
		return p - (p * disc / 100)
	}
	return p
}
