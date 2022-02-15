package main

import (
	"fmt"
	accounts "github.com/oo/bank/accounts"
)

func main() {
	ccAccount := accounts.CheckingAccount{
		Owner:   "Breno",
		Balance: 300,
	}

	ccAccount2 := accounts.CheckingAccount{
		Owner:   "Teste2",
		Balance: 100,
	}

	status := ccAccount.Transfer(200., &ccAccount2)
	fmt.Println(status)
	fmt.Println(ccAccount, ccAccount2)
}
