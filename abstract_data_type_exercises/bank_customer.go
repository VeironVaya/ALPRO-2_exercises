package main

import "fmt"

type Customer struct {
	customerID    string
	customerName  string
	bankName      string
	accountNumber string
}

type ArrayCustomer [2022]Customer

func inputCustomerData(data *ArrayCustomer, n *int) {
	var i int
	for i = 0; i < *n; i++ {
		fmt.Print("Customer ID: ")
		fmt.Scan(&data[i].customerID)
		fmt.Print("Customer Name: ")
		fmt.Scan(&data[i].customerName)
		fmt.Print("Bank Name: ")
		fmt.Scan(&data[i].bankName)
		fmt.Print("Account Number: ")
		fmt.Scan(&data[i].accountNumber)
	}
}

func printAllCustomers(data ArrayCustomer, n int) {
	var i int
	for i = 0; i < n; i++ {
		if data[i].bankName == "MANDIRI" {
			fmt.Println("\nCustomer", i+1)
			fmt.Println("ID:", data[i].customerID)
			fmt.Println("Name:", data[i].customerName)
			fmt.Println("Bank:", data[i].bankName)
			fmt.Println("Account:", data[i].accountNumber)
		}
	}
}

func main() {
	var customers ArrayCustomer
	var total int = 10

	inputCustomerData(&customers, &total)
	printAllCustomers(customers, total)
}
