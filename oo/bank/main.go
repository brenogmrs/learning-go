package main

import "fmt"

type CheckingAccount struct {
	titular 	  string
	agencyNumber  int
	accountNumber int
	balance 	  float64
}



func main() {
	brenosAccount := CheckingAccount{
		"Breno",
		234,
		123456,
		125.50,
	}

	var brenoAccount *CheckingAccount
	brenoAccount = new(CheckingAccount)
	brenoAccount.titular = "Gabriela"
	brenoAccount.balance = 500

	fmt.Println(*brenoAccount)
}