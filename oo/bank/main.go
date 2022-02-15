package main

import (
	"fmt"

	"github.com/oo/bank/accounts"
	"github.com/oo/bank/customers"
)

func main() {

	customer1 := customers.AccountOwner{
		Name:       "Breno",
		Cpf:        "123123",
		Occupation: "it",
	}

	ccAccount := accounts.CheckingAccount{
		Owner: customer1,
	}

	ccAccount.Deposit(200)

	fmt.Println(ccAccount.GetBalance())

}
