package main

import (
	"fmt"

	"github.com/oo/bank/accounts"
	"github.com/oo/bank/customers"
)

type verifyAccount interface {
	Withdraw(float64) string
}

func payBoleto(account verifyAccount, BoletoValue float64) {
	account.Withdraw(BoletoValue)
}

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

	ccAccount.Deposit(200)

	payBoleto(&ccAccount, 60)

	fmt.Println(ccAccount.GetBalance())

	scAccount.Deposit(200)

	payBoleto(&scAccount, 60)

	fmt.Println(scAccount.GetBalance())

}
