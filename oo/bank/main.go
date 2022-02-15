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

	customer2 := customers.AccountOwner{
		Name:       "Teste2",
		Cpf:        "123123",
		Occupation: "ti",
	}

	ccAccount := accounts.CheckingAccount{
		Owner: customer1,
	}

	ccAccount2 := accounts.CheckingAccount{
		Owner: customer2,
	}

	status := ccAccount.Transfer(200., &ccAccount2)
	fmt.Println(status)
	fmt.Println(ccAccount, ccAccount2)
}
