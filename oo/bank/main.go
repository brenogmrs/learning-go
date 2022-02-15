package main

import (
	"fmt"
)

func main() {
	ccAccount := CheckingAccount{
		Owner:   "Breno",
		Balance: 300,
	}

	ccAccount2 := CheckingAccount{
		Owner:   "Teste2",
		Balance: 100,
	}

	status := ccAccount.Transfer(200., &ccAccount2)
	fmt.Println(status)
	fmt.Println(ccAccount, ccAccount2)
}
