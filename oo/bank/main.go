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

	scAccount := accounts.SavingsAccount{
		Owner:         customer1,
		AgencyNumber:  123,
		AccountNumber: 123123,
		Operation:     1,
	}

	customer2 := customers.AccountOwner{
		Name:       "Teste2",
		Cpf:        "123123",
		Occupation: "ti",
	}

	ccAccount := accounts.CheckingAccount{
		Owner: customer2,
	}

	fmt.Println(scAccount, ccAccount)

}
